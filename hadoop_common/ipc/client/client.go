package ipc

import (
	"bytes"
	"code.google.com/p/goprotobuf/proto"
	"errors"
	"fmt"
	"github.com/hortonworks/gohadoop"
	"github.com/hortonworks/gohadoop/hadoop_common"
	"github.com/nu7hatch/gouuid"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
)

type Client struct {
	ClientId      *uuid.UUID
	Ugi           *hadoop_common.UserInformationProto
	ServerAddress string
	TCPNoDelay    bool
}

type connection struct {
	con *net.TCPConn
}

type connection_id struct {
	user     string
	protocol string
}

type call struct {
	callId     int32
	procedure  proto.Message
	request    proto.Message
	response   proto.Message
	err        *error
	retryCount int32
}

func (c *Client) String() string {
	buf := bytes.NewBufferString("")
	fmt.Fprint(buf, "<clientId:", c.ClientId)
	fmt.Fprint(buf, ", server:", c.ServerAddress)
	fmt.Fprint(buf, ">")
	return buf.String()
}

func (c *Client) Call(rpc *hadoop_common.RequestHeaderProto, rpcRequest proto.Message, rpcResponse proto.Message) error {
	// Create connection_id
	connectionId := connection_id{user: *c.Ugi.RealUser, protocol: *rpc.DeclaringClassProtocolName}

	// Get connection to server
	//log.Println("Connecting...", c)
	conn, err := getConnection(c, &connectionId)
	if err != nil {
		return err
	}
	// TODO return connection to pool?

	log.Printf("request: %T: %+v", rpcRequest, rpcRequest)

	// Create call and send request
	rpcCall := call{callId: 0, procedure: rpc, request: rpcRequest, response: rpcResponse}
	err = sendRequest(c, conn, &rpcCall)
	if err != nil {
		log.Fatal("sendRequest", err)
		return err
	}

	// Read & return response
	err = c.readResponse(conn, &rpcCall)

	log.Printf("response: %T: %+v", rpcResponse, rpcResponse)

	return err
}

var connectionPool = struct {
	sync.RWMutex
	connections map[connection_id]*connection
}{connections: make(map[connection_id]*connection)}

func getConnection(c *Client, connectionId *connection_id) (*connection, error) {
	// Try to re-use an existing connection
	connectionPool.RLock()
	con := connectionPool.connections[*connectionId]
	connectionPool.RUnlock()

	// If necessary, create a new connection and save it in the connection-pool
	var err error
	if con == nil {
		log.Printf("SETTING UP NEW CONNECTION")

		con, err = setupConnection(c)
		if err != nil {
			log.Fatal("Couldn't setup connection: ", err)
			return nil, err
		}

		connectionPool.Lock()
		connectionPool.connections[*connectionId] = con
		connectionPool.Unlock()

		err = writeConnectionHeader(con)
		if err != nil {
			// TODO what happens to the con?
			return nil, err
		}

		log.Printf("SETTING UP")
		if false {
			err = setupSaslConnection(c, con, connectionId)
			if err != nil {
				// TODO what happens to the con?
				return nil, err
			}
		}
		log.Printf("SETTING UP DONE")

		err = writeConnectionContext(c, con, connectionId)
		if err != nil {
			// TODO what happens to the con?
			return nil, err
		}
	}

	return con, nil
}

func setupConnection(c *Client) (*connection, error) {
	addr, _ := net.ResolveTCPAddr("tcp", c.ServerAddress)
	log.Printf("dialing %v", addr)
	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Printf("error: %v", err)
		return nil, err
	} else {
		log.Printf("Successfully connected %v", c)
	}

	// TODO: Ping thread

	// Set tcp no-delay
	tcpConn.SetNoDelay(c.TCPNoDelay)

	return &connection{tcpConn}, nil
}

func writeConnectionHeader(conn *connection) error {
	// RPC_HEADER
	if _, err := conn.con.Write(gohadoop.RPC_HEADER); err != nil {
		log.Fatal("conn.Write gohadoop.RPC_HEADER", err)
		return err
	}

	// RPC_VERSION
	if _, err := conn.con.Write(gohadoop.VERSION); err != nil {
		log.Fatal("conn.Write gohadoop.VERSION", err)
		return err
	}

	// RPC_SERVICE_CLASS
	if serviceClass, err := gohadoop.ConvertFixedToBytes(gohadoop.RPC_SERVICE_CLASS); err != nil {
		log.Fatal("binary.Write", err)
		return err
	} else if _, err := conn.con.Write(serviceClass); err != nil {
		log.Fatal("conn.Write RPC_SERVICE_CLASS", err)
		return err
	}

	// AuthProtocol
	//if authProtocol, err := gohadoop.ConvertFixedToBytes(gohadoop.AUTH_PROTOCOL_SASL); err != nil {
	if authProtocol, err := gohadoop.ConvertFixedToBytes(gohadoop.AUTH_PROTOCOL_NONE); err != nil {
		log.Fatal("WTF AUTH_PROTOCOL_NONE", err)
		return err
	} else if _, err := conn.con.Write(authProtocol); err != nil {
		log.Fatal("conn.Write gohadoop.AUTH_PROTOCOL_SASL", err)
		return err
	}

	return nil
}

// https://github.com/youtube/vitess/tree/master/go/rpcwrap/auth

// list parser:
// https://github.com/golang/gddo/blob/master/httputil/header/header.go

// ParseList parses a comma separated list of values. Commas are ignored in
// quoted strings. Quoted values are not unescaped or unquoted. Whitespace is
// trimmed.
func ParseList(s string) map[string]string {
	var kvs []string

	begin := 0
	end := 0
	escape := false
	quote := false
	for i := 0; i < len(s); i++ {
		b := s[i]
		switch {
		case escape:
			escape = false
			end = i + 1
		case quote:
			switch b {
			case '\\':
				escape = true
			case '"':
				quote = false
			}
			end = i + 1
		case b == '"':
			quote = true
			end = i + 1
		case octetTypes[b]&isSpace != 0:
			if begin == end {
				begin = i + 1
				end = begin
			}
		case b == ',':
			if begin < end {
				kvs = append(kvs, s[begin:end])
			}
			begin = i + 1
			end = begin
		default:
			end = i + 1
		}
	}
	if begin < end {
		kvs = append(kvs, s[begin:end])
	}

	results := make(map[string]string, len(kvs))

	for _, kv := range kvs {
		parts := strings.SplitN(kv, "=", 2)
		k := parts[0]
		if len(parts) == 1 {
			results[k] = ""
			continue
		}
		results[k] = parts[1]
	}

	log.Printf("results: %v", results)

	return results
}

// Octet types from RFC 2616.
var octetTypes [256]octetType

type octetType byte

const (
	isToken octetType = 1 << iota
	isSpace
)

func init() {
	// OCTET      = <any 8-bit sequence of data>
	// CHAR       = <any US-ASCII character (octets 0 - 127)>
	// CTL        = <any US-ASCII control character (octets 0 - 31) and DEL (127)>
	// CR         = <US-ASCII CR, carriage return (13)>
	// LF         = <US-ASCII LF, linefeed (10)>
	// SP         = <US-ASCII SP, space (32)>
	// HT         = <US-ASCII HT, horizontal-tab (9)>
	// <">        = <US-ASCII double-quote mark (34)>
	// CRLF       = CR LF
	// LWS        = [CRLF] 1*( SP | HT )
	// TEXT       = <any OCTET except CTLs, but including LWS>
	// separators = "(" | ")" | "<" | ">" | "@" | "," | ";" | ":" | "\" | <"> | "/" | "[" | "]" | "?" | "=" | "{" | "}" | SP | HT
	// token      = 1*<any CHAR except CTLs or separators>
	// qdtext     = <any TEXT except <">>

	for c := 0; c < 256; c++ {
		var t octetType
		isCtl := c <= 31 || c == 127
		isChar := 0 <= c && c <= 127
		isSeparator := strings.IndexRune(" \t\"(),/:;<=>?@[]\\{}", rune(c)) >= 0
		if strings.IndexRune(" \t\r\n", rune(c)) >= 0 {
			t |= isSpace
		}
		if isChar && !isCtl && !isSeparator {
			t |= isToken
		}
		octetTypes[c] = t
	}
}

func setupSaslConnection(c *Client, conn *connection, connectionId *connection_id) error {
	err := writeSaslNegotiate(c, conn, connectionId)
	if err != nil {
		return err
	}

	var saslMessage hadoop_common.RpcSaslProto
	err = c.readResponseSasl(conn, &saslMessage)
	if err != nil {
		return err
	}

	log.Printf("response: %+v", saslMessage)

	switch saslMessage.GetState() {
	case hadoop_common.RpcSaslProto_NEGOTIATE:
		for _, auth := range saslMessage.GetAuths() {
			log.Printf("auth: %+v", auth)
			switch auth.GetMethod() {
			case "SIMPLE":
				log.Printf("simple, next please")
			case "TOKEN":
				switch auth.GetMechanism() {
				case "DIGEST-MD5":
					evaluateChallenge(auth.Challenge)
				default:
					panic("not implemented")
				}
			default:
				panic("not implemented")
			}
		}
	default:
		panic("not implemented")
	}

	panic("splat")
}

// realm=\"default\",nonce=\"VpkdRl6/YsgqZfaejtJDM/oUhD7PrlUjCTkAA0kB\",qop=\"auth\",charset=utf-8,algorithm=md5-sess

func evaluateChallenge(challenge []byte) {
	parsed := ParseList(string(challenge))
	cleaned := make(map[string]string, len(parsed))
	for k, v := range parsed {
		w, err := strconv.Unquote(v)
		if err != nil {
			w = v
		}
		cleaned[strings.ToLower(k)] = w
	}

	if strings.ToLower(cleaned["algorithm"]) != "md5-sess" {
		panic("splat")
	}
	nonce := cleaned["nonce"]
	if len(nonce) == 0 {
		panic("splat")
	}

	panic("not implemented")
}

func writeSaslNegotiate(c *Client, conn *connection, connectionId *connection_id) error {
	var callId int32 = -33
	rpcReqHeaderProto := hadoop_common.RpcRequestHeaderProto{
		RpcKind:    &gohadoop.RPC_PROTOCOL_BUFFFER,
		RpcOp:      &gohadoop.RPC_FINAL_PACKET,
		CallId:     &callId,
		ClientId:   []byte{},
		RetryCount: &gohadoop.RPC_DEFAULT_RETRY_COUNT,
	}

	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		return err
	}

	saslState := hadoop_common.RpcSaslProto_NEGOTIATE
	negotiateRequest := hadoop_common.RpcSaslProto{
		State: &saslState,
	}

	negReqBytes, err := proto.Marshal(&negotiateRequest)
	if err != nil {
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(negReqBytes) + sizeVarint(len(negReqBytes))
	var tLen int32 = int32(totalLength)
	totalLengthBytes, err := gohadoop.ConvertFixedToBytes(tLen)
	if err != nil {
		return err
	}
	if _, err := conn.con.Write(totalLengthBytes); err != nil {
		return err
	}

	if err := writeDelimitedBytes(conn, rpcReqHeaderProtoBytes); err != nil {
		return err
	}
	if err := writeDelimitedBytes(conn, negReqBytes); err != nil {
		return err
	}

	return nil
}

func writeConnectionContext(c *Client, conn *connection, connectionId *connection_id) error {
	// Create hadoop_common.IpcConnectionContextProto
	ugi, _ := gohadoop.CreateSimpleUGIProto()
	ipcCtxProto := hadoop_common.IpcConnectionContextProto{UserInfo: ugi, Protocol: &connectionId.protocol}

	// Create RpcRequestHeaderProto
	var callId int32 = -3
	var clientId [16]byte = [16]byte(*c.ClientId)
	rpcReqHeaderProto := hadoop_common.RpcRequestHeaderProto{RpcKind: &gohadoop.RPC_PROTOCOL_BUFFFER, RpcOp: &gohadoop.RPC_FINAL_PACKET, CallId: &callId, ClientId: clientId[0:16], RetryCount: &gohadoop.RPC_DEFAULT_RETRY_COUNT}

	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Fatal("proto.Marshal(&rpcReqHeaderProto)", err)
		return err
	}

	ipcCtxProtoBytes, _ := proto.Marshal(&ipcCtxProto)
	if err != nil {
		log.Fatal("proto.Marshal(&ipcCtxProto)", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(ipcCtxProtoBytes) + sizeVarint(len(ipcCtxProtoBytes))
	var tLen int32 = int32(totalLength)
	if totalLengthBytes, err := gohadoop.ConvertFixedToBytes(tLen); err != nil {
		log.Fatal("ConvertFixedToBytes(totalLength)", err)
		return err
	} else if _, err := conn.con.Write(totalLengthBytes); err != nil {
		log.Fatal("conn.con.Write(totalLengthBytes)", err)
		return err
	}

	if err := writeDelimitedBytes(conn, rpcReqHeaderProtoBytes); err != nil {
		log.Fatal("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, ipcCtxProtoBytes); err != nil {
		log.Fatal("writeDelimitedBytes(conn, ipcCtxProtoBytes)", err)
		return err
	}

	return nil
}

func sizeVarint(x int) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func sendRequest(c *Client, conn *connection, rpcCall *call) error {
	//log.Println("About to call RPC: ", rpcCall.procedure)

	// 0. RpcRequestHeaderProto
	var clientId [16]byte = [16]byte(*c.ClientId)
	rpcReqHeaderProto := hadoop_common.RpcRequestHeaderProto{RpcKind: &gohadoop.RPC_PROTOCOL_BUFFFER, RpcOp: &gohadoop.RPC_FINAL_PACKET, CallId: &rpcCall.callId, ClientId: clientId[0:16], RetryCount: &rpcCall.retryCount}
	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Fatal("proto.Marshal(&rpcReqHeaderProto)", err)
		return err
	}

	// 1. RequestHeaderProto
	requestHeaderProto := rpcCall.procedure
	requestHeaderProtoBytes, err := proto.Marshal(requestHeaderProto)
	if err != nil {
		log.Fatal("proto.Marshal(&requestHeaderProto)", err)
		return err
	}

	// 2. Param
	paramProto := rpcCall.request
	paramProtoBytes, err := proto.Marshal(paramProto)
	if err != nil {
		log.Fatal("proto.Marshal(&paramProto)", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(requestHeaderProtoBytes) + sizeVarint(len(requestHeaderProtoBytes)) + len(paramProtoBytes) + sizeVarint(len(paramProtoBytes))
	var tLen int32 = int32(totalLength)
	if totalLengthBytes, err := gohadoop.ConvertFixedToBytes(tLen); err != nil {
		log.Fatal("ConvertFixedToBytes(totalLength)", err)
		return err
	} else {
		if _, err := conn.con.Write(totalLengthBytes); err != nil {
			log.Fatal("conn.con.Write(totalLengthBytes)", err)
			return err
		}
	}

	if err := writeDelimitedBytes(conn, rpcReqHeaderProtoBytes); err != nil {
		log.Fatal("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, requestHeaderProtoBytes); err != nil {
		log.Fatal("writeDelimitedBytes(conn, requestHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, paramProtoBytes); err != nil {
		log.Fatal("writeDelimitedBytes(conn, paramProtoBytes)", err)
		return err
	}

	//log.Println("Succesfully sent request of length: ", totalLength)

	return nil
}

func writeDelimitedTo(conn *connection, msg proto.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("proto.Marshal(msg)", err)
		return err
	}
	return writeDelimitedBytes(conn, msgBytes)
}

func writeDelimitedBytes(conn *connection, data []byte) error {
	if _, err := conn.con.Write(proto.EncodeVarint(uint64(len(data)))); err != nil {
		log.Fatal("conn.con.Write(proto.EncodeVarint(uint64(len(data))))", err)
		return err
	}
	if _, err := conn.con.Write(data); err != nil {
		log.Fatal("conn.con.Write(data)", err)
		return err
	}

	return nil
}

func (c *Client) readResponse(conn *connection, rpcCall *call) error {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte
	if _, err := conn.con.Read(totalLengthBytes[0:4]); err != nil {
		log.Fatal("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	if err := gohadoop.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Fatal("gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength)", err)
		return err
	}

	log.Printf("totalLength=%d", totalLength)

	var responseBytes []byte = make([]byte, totalLength)
	if _, err := conn.con.Read(responseBytes); err != nil {
		log.Fatal("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := hadoop_common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Fatal("readDelimited(responseBytes, rpcResponseHeaderProto)", err)
		return err
	}
	log.Printf("Received rpcResponseHeaderProto = %+v", rpcResponseHeaderProto)

	err = c.checkRpcHeader(&rpcResponseHeaderProto)
	if err != nil {
		log.Fatal("c.checkRpcHeader failed", err)
		return err
	}

	if *rpcResponseHeaderProto.Status == hadoop_common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		_, err = readDelimited(responseBytes[off:], rpcCall.response)
	} else {
		log.Println("RPC failed with status: ", rpcResponseHeaderProto.Status.String())
		errorDetails := [4]string{rpcResponseHeaderProto.Status.String(), "ServerDidNotSetExceptionClassName", "ServerDidNotSetErrorMsg", "ServerDidNotSetErrorDetail"}
		if rpcResponseHeaderProto.ExceptionClassName != nil {
			errorDetails[0] = *rpcResponseHeaderProto.ExceptionClassName
		}
		if rpcResponseHeaderProto.ErrorMsg != nil {
			errorDetails[1] = *rpcResponseHeaderProto.ErrorMsg
		}
		if rpcResponseHeaderProto.ErrorDetail != nil {
			errorDetails[2] = rpcResponseHeaderProto.ErrorDetail.String()
		}
		err = errors.New(strings.Join(errorDetails[:], ":"))
	}
	return err
}

func (c *Client) readResponseSasl(conn *connection, saslMessage *hadoop_common.RpcSaslProto) error {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte
	if _, err := conn.con.Read(totalLengthBytes[0:4]); err != nil {
		log.Fatal("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	if err := gohadoop.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Fatal("gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength)", err)
		return err
	}

	log.Printf("totalLength=%d", totalLength)

	var responseBytes []byte = make([]byte, totalLength)
	if _, err := conn.con.Read(responseBytes); err != nil {
		log.Fatal("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := hadoop_common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Fatal("readDelimited(responseBytes, rpcResponseHeaderProto)", err)
		return err
	}
	log.Printf("Received rpcResponseHeaderProto = %+v", rpcResponseHeaderProto)

	err = c.checkRpcHeader(&rpcResponseHeaderProto)
	if err != nil {
		log.Fatal("c.checkRpcHeader failed", err)
		return err
	}

	if *rpcResponseHeaderProto.Status == hadoop_common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		_, err = readDelimited(responseBytes[off:], saslMessage)
	} else {
		log.Println("RPC failed with status: ", rpcResponseHeaderProto.Status.String())
		errorDetails := [4]string{rpcResponseHeaderProto.Status.String(), "ServerDidNotSetExceptionClassName", "ServerDidNotSetErrorMsg", "ServerDidNotSetErrorDetail"}
		if rpcResponseHeaderProto.ExceptionClassName != nil {
			errorDetails[0] = *rpcResponseHeaderProto.ExceptionClassName
		}
		if rpcResponseHeaderProto.ErrorMsg != nil {
			errorDetails[1] = *rpcResponseHeaderProto.ErrorMsg
		}
		if rpcResponseHeaderProto.ErrorDetail != nil {
			errorDetails[2] = rpcResponseHeaderProto.ErrorDetail.String()
		}
		err = errors.New(strings.Join(errorDetails[:], ":"))
	}
	return err
}

func readDelimited(rawData []byte, msg proto.Message) (int, error) {
	headerLength, off := proto.DecodeVarint(rawData)
	if off == 0 {
		log.Fatal("proto.DecodeVarint(rawData) returned zero")
		return -1, nil
	}
	err := proto.Unmarshal(rawData[off:off+int(headerLength)], msg)
	if err != nil {
		log.Fatal("proto.Unmarshal(rawData[off:off+headerLength]) ", err)
		return -1, err
	}

	return off + int(headerLength), nil
}

func (c *Client) checkRpcHeader(rpcResponseHeaderProto *hadoop_common.RpcResponseHeaderProto) error {
	var callClientId [16]byte = [16]byte(*c.ClientId)
	var headerClientId []byte = []byte(rpcResponseHeaderProto.ClientId)
	if len(rpcResponseHeaderProto.ClientId) > 0 {
		log.Printf("len(headerClientId)=%d", len(headerClientId))
		if !bytes.Equal(callClientId[0:16], headerClientId[0:16]) {
			log.Fatal("Incorrect clientId: ", headerClientId)
			return errors.New("Incorrect clientId")
		}
	}
	return nil
}

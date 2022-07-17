package util

import (
	"errors"
	"io"

	"bitwormhole.com/starter/vlog"
)

// Close 关闭指定的 closer
func Close(c io.Closer) error {
	if c == nil {
		return nil
	}
	err := c.Close()
	if err != nil {
		vlog.Warn(err)
	}
	return err
}

// PumpStream 把数据从 src 读出来，并写入到 dst. 返回成功写入的字节数.
func PumpStream(src io.Reader, dst io.Writer, buffer []byte) (int64, error) {
	if buffer == nil {
		// 默认缓冲区大小为 4k
		buffer = make([]byte, 1024*4)
	}
	var count int64 = 0
	for {
		cb1, err1 := src.Read(buffer)
		if cb1 > 0 {
			cb2, err2 := dst.Write(buffer[0:cb1])
			if err2 != nil {
				return count, err2
			}
			if cb1 != cb2 {
				return count, errors.New("count.bytes.in != count.bytes.out")
			}
			count += int64(cb2)
		}
		if err1 != nil {
			if IsEOF(err1) {
				return count, nil
			}
			return count, err1
		}
	}
}

// IsEOF 检查并判断 err 是否为 EOF
func IsEOF(err error) bool {
	return err == io.EOF
}

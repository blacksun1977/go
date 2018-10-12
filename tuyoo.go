package jsoniter

// 2018-10-08
// 特殊的直接将结果写入到文件当中，避免二次的内存拷贝
func (cfg *frozenConfig) Marshal2File(lineDate int, writefun func(lineDate int, p []byte) error, v interface{}) error {
	stream := cfg.BorrowStream(nil)
	defer cfg.ReturnStream(stream)
	stream.WriteVal(v)
	if stream.Error != nil {
		return stream.Error
	}
	// 添加换行符号
	stream.writeByte('\n')
	return writefun(lineDate, stream.Buffer())
}

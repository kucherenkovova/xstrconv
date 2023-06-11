package xstrconv

import "strconv"

func ParseInt8(s string, base int) (int8, error) {
	result, err := strconv.ParseInt(s, base, 8)
	if err != nil {
		return 0, err
	}

	return int8(result), nil
}

func ParseInt16(s string, base int) (int16, error) {
	result, err := strconv.ParseInt(s, base, 16)
	if err != nil {
		return 0, err
	}

	return int16(result), nil
}

func ParseInt32(s string, base int) (int32, error) {
	result, err := strconv.ParseInt(s, base, 32)
	if err != nil {
		return 0, err
	}

	return int32(result), nil
}

func ParseInt64(s string, base int) (int64, error) {
	return strconv.ParseInt(s, base, 64)
}

func ParseDecimalInt8(s string) (int8, error) {
	return ParseInt8(s, 10)
}

func ParseDecimalInt16(s string) (int16, error) {
	return ParseInt16(s, 10)
}

func ParseDecimalInt32(s string) (int32, error) {
	return ParseInt32(s, 10)
}

func ParseDecimalInt64(s string) (int64, error) {
	return ParseInt64(s, 10)
}

func ParseUint8(s string, base int) (uint8, error) {
	result, err := strconv.ParseUint(s, base, 8)
	if err != nil {
		return 0, err
	}

	return uint8(result), nil
}

func ParseUint16(s string, base int) (uint16, error) {
	result, err := strconv.ParseUint(s, base, 16)
	if err != nil {
		return 0, err
	}

	return uint16(result), nil
}

func ParseUint32(s string, base int) (uint32, error) {
	result, err := strconv.ParseUint(s, base, 32)
	if err != nil {
		return 0, err
	}

	return uint32(result), nil
}

func ParseUint64(s string, base int) (uint64, error) {
	return strconv.ParseUint(s, base, 64)
}

func ParseDecimalUint8(s string) (uint8, error) {
	return ParseUint8(s, 10)
}

func ParseDecimalUint16(s string) (uint16, error) {
	return ParseUint16(s, 10)
}

func ParseDecimalUint32(s string) (uint32, error) {
	return ParseUint32(s, 10)
}

func ParseDecimalUint64(s string) (uint64, error) {
	return ParseUint64(s, 10)
}

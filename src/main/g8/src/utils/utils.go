/*
Package utils provides a utility/helper functions that can be used by the application.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@since template-v0.1.0
*/
package utils

import (
    "bytes"
    olaf2 "github.com/btnguyen2k/consu/olaf"
    "math/rand"
    "net"
    "strconv"
    "strings"
    "time"
)

// global variables
var (
    // Location should be initialized during application bootstrap
    Location *time.Location
)

func getMacAddr() string {
    interfaces, err := net.Interfaces()
    if err == nil {
        for _, i := range interfaces {
            if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
                // Don't use random as we have a real address
                return i.HardwareAddr.String()
            }
        }
    }
    return ""
}

func getMacAddrAsLong() int64 {
    mac, _ := strconv.ParseInt(strings.Replace(getMacAddr(), ":", "", -1), 16, 64)
    return mac
}

var olaf = olaf2.NewOlaf(getMacAddrAsLong())

// UniqueId generates a unique id as hex-string
func UniqueId() string {
    return olaf.Id128Hex()
}

// UniqueIdSmall generates a unique id as hex-string
func UniqueIdSmall() string {
    return olaf.Id64Ascii()
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

/*
RandomString generates a random string with specified length.
*/
func RandomString(l int) string {
    b := make([]byte, l)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

// /*
// PadRight adds "0" right right of a string until its length reach a specific value.
// */
// func PadRight(str string, l int) string {
//     for len(str) < l {
//         str += "0"
//     }
//     return str
// }
//
// /*
// AesEncrypt encrypts a block of data using AES/CTR mode.
//
// IV is put at the beginning of the cipher data.
// */
// func AesEncrypt(key, data []byte) ([]byte, error) {
//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return nil, err
//     }
//     iv := []byte(PadRight(strconv.FormatInt(time.Now().UnixNano(), 16), 16))
//     cipherData := make([]byte, 16+len(data))
//     copy(cipherData, iv)
//     ctr := cipher.NewCTR(block, iv)
//     ctr.XORKeyStream(cipherData[16:], data)
//     return cipherData, nil
// }
//
// /*
// AesDecrypt decrypts a block of encrypted data using AES/CTR mode.
//
// Assuming IV is put at the beginning of the cipher data.
// */
// func AesDecrypt(key, encryptedData []byte) ([]byte, error) {
//     block, err := aes.NewCipher(key)
//     if err != nil {
//         return nil, err
//     }
//     iv := encryptedData[0:16]
//     data := make([]byte, len(encryptedData)-16)
//     ctr := cipher.NewCTR(block, iv)
//     ctr.XORKeyStream(data, encryptedData[16:])
//     return data, nil
// }
//
// /*
// ZlibCompress compresses data using zlib.
// */
// func ZlibCompress(data []byte) []byte {
//     var b bytes.Buffer
//     w, _ := zlib.NewWriterLevel(&b, zlib.BestCompression)
//     w.Write(data)
//     w.Close()
//     return b.Bytes()
// }
//
// /*
// ZlibDecompress decompressed compressed-data using zlib.
// */
// func ZlibDecompress(compressedData []byte) ([]byte, error) {
//     r, err := zlib.NewReader(bytes.NewReader(compressedData))
//     if err != nil {
//         return nil, err
//     }
//     var b bytes.Buffer
//     _, err = io.Copy(&b, r)
//     r.Close()
//     return b.Bytes(), err
// }
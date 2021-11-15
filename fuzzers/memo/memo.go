package memofuzz

import memo "gitlab.com/thorchain/thornode/x/thorchain/memo"

func Fuzz(data []byte) int {
    memo.ParseMemo(string(data))
    return 0
}

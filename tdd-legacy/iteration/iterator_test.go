package iteration

import "testing"

func TestRepeat(t *testing.T) {

    repeated := Repeat("a")
    expected := "aaaaa"

    if repeated != expected {
        t.Errorf("expected %q but got %q", expected, repeated)
    }

}

func BenchmarkRepeat(b *testing.B) {
    // for b.Loop(){}
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}

func BenchmarkRepeats(b *testing.B) {

    implementations := []struct {
        name string
        fn   func(string) string
    }{
        {"RepeatV1", RepeatV1},
        {"RepeatV2", RepeatV2},
        {"Repeat", Repeat},
    }

    testInput := "a"

    // 用于防止编译器优化
    var globalResult string

    for _, impl := range implementations {
        b.Run(impl.name, func(b *testing.B) {
            var localResult string
            b.ReportAllocs()
            b.ResetTimer()
            for range b.N {
                // 赋值给局部变量，防止优化
                localResult = impl.fn(testInput)
            }
            globalResult = localResult
        })

    }
    _ = globalResult

}

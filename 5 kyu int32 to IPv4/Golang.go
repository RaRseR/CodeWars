import ("fmt")

func Int32ToIp(n uint32) string {
    return fmt.Sprintf("%d.%d.%d.%d", (n >> 24) % 256, (n >> 16) % 256, (n >> 8) % 256, n % 256)
}

Loadtesting done via k6

run using following command:
`docker run -i loadimpact/k6 run -d 60s -u 10 - <test.js`

During testing, found race condition:
```
wordgame-server-wordgame-1       | fatal error: concurrent map writes
wordgame-server-wordgame-1       |
wordgame-server-wordgame-1       | goroutine 195 [running]:
wordgame-server-wordgame-1       | runtime.throw({0x76f671, 0x7f63a16dde10})
wordgame-server-wordgame-1       | 	/usr/local/go/src/runtime/panic.go:1198 +0x71 fp=0xc0006ca798 sp=0xc0006ca768 pc=0x435111
wordgame-server-wordgame-1       | runtime.mapassign_faststr(0xc0006ca800, 0xc0011d03de, {0xc0009361b0, 0x24})
wordgame-server-wordgame-1       | 	/usr/local/go/src/runtime/map_faststr.go:211 +0x39c fp=0xc0006ca800 sp=0xc0006ca798 pc=0x412c3c
wordgame-server-wordgame-1       | main.(*MemoryStore).SaveGame(0x7d84c8, {{0xc0009361b0, 0x24}, {0xc000721e84, 0x2}, {0xc0011d03de, 0x2}, 0x6})
wordgame-server-wordgame-1       | 	/build/store.go:16 +0x54 fp=0xc0006ca830 sp=0xc0006ca800 pc=0x6d3db4
wordgame-server-wordgame-1       | main.(*player).NewGame(0xc000118060)
wordgame-server-wordgame-1       | 	/build/player.go:64 +0x285 fp=0xc0006ca948 sp=0xc0006ca830 pc=0x6d2a65
wordgame-server-wordgame-1       | main.(*server).NewGame(0xc0006caa30, {0x7e2508, 0xc000b7c620}, 0x0)
wordgame-server-wordgame-1       | 	/build/server.go:24 +0x2b9 fp=0xc0006caa18 sp=0xc0006ca948 pc=0x6d3259
wordgame-server-wordgame-1       | main.(*server).NewGame-fm({0x7e2508, 0xc000b7c620}, 0x0)
wordgame-server-wordgame-1       | 	/build/server.go:15 +0x3c fp=0xc0006caa48 sp=0xc0006caa18 pc=0x6d46fc
wordgame-server-wordgame-1       | net/http.HandlerFunc.ServeHTTP(0x0, {0x7e2508, 0xc000b7c620}, 0x0)
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:2047 +0x2f fp=0xc0006caa70 sp=0xc0006caa48 pc=0x61cb0f
wordgame-server-wordgame-1       | net/http.(*ServeMux).ServeHTTP(0x0, {0x7e2508, 0xc000b7c620}, 0xc00065e600)
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:2425 +0x149 fp=0xc0006caac0 sp=0xc0006caa70 pc=0x61e409
wordgame-server-wordgame-1       | net/http.serverHandler.ServeHTTP({0xc0006cc330}, {0x7e2508, 0xc000b7c620}, 0xc00065e600)
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:2879 +0x43b fp=0xc0006cab80 sp=0xc0006caac0 pc=0x61f73b
wordgame-server-wordgame-1       | net/http.(*conn).serve(0xc0006de140, {0x7e3ba0, 0xc000118210})
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:1930 +0xb08 fp=0xc0006cafb8 sp=0xc0006cab80 pc=0x61be68
wordgame-server-wordgame-1       | net/http.(*Server).Serve·dwrap·87()
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:3034 +0x2e fp=0xc0006cafe0 sp=0xc0006cafb8 pc=0x62008e
wordgame-server-wordgame-1       | runtime.goexit()
wordgame-server-wordgame-1       | 	/usr/local/go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc0006cafe8 sp=0xc0006cafe0 pc=0x4655a1
wordgame-server-wordgame-1       | created by net/http.(*Server).Serve
wordgame-server-wordgame-1       | 	/usr/local/go/src/net/http/server.go:3034 +0x4e8
wordgame-server-wordgame-1       |
```

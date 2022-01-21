## Performance Debugging

So basically immediately I found that the issue was with SSL, but that was basically given from the customer's problem statement.
What I found strange was that it was spending quite a bit of time doing exponents of large integers. With how SSL works, this is independant of the data being sent/recieved, but is dependant on the size of the rsa key being used. Its almost as if they are using a really big rsa key (beyond 4096). This is also supported by it taking a lot more memory, since a larger key would need more memory, but the difference in memory requirements doesn't fit with it just being a larger key.
I would say that this is most likely an issue with the cert on either the server's or client's side since its happening before the handshake completes and is part of the standard library, which should be efficient under most circumstances. A really big certificate (possibly chain) on either the client (if mutual TLS) or server side seems to be the most likely cause.
see https://github.com/golang/go/blob/go1.17.5/src/crypto/tls/handshake_server_tls13.go#L622
end 30 mins, moving on

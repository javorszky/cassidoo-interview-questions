# 12th July 2021

> This week’s question:
>
> Given an IPv4 address and a netmask in CIDR notation, return a boolean specifying whether the IP address is inside the given range.
>
>Example:
> ```bash
> $ inRange("192.168.4.27", "192.168.0.0/16")
> $ true
> 
> $ inRange("192.168.4.27", "192.168.1.0/24")
> $ false
> ```

## what I've done
There are two solutions. One of them, the lazy one, is using go's built in `net` package. That one has a `parseCIDR` function that returns an IPNet pointer, among other things, and that IPNet has a `Contains` method you can pass an IP to.

It also has a `parseIP` method that returns an IP struct to be used in the Contains method.

The other one is hand parsing both of them. I wrote regexes for both the IP and CIDR notations, I pick the values, convert them to bytes after some rudimentary error checking, turn the cidr range into a bitmask, xor the ip and the bitmask, and return whether the ip falls inside the mask (xor value is less than the mask), or not.

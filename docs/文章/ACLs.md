# ACLs 访问控制权限

## Specifications

```
a)
access-list 1 deny 170.16.40.0 0.0.0.255
access-list 1 permit 170.16.0.0 0.0.255.255
interface Ethernet1/0
ip access-group 1 in

b)
access-list 2 deny ip host 170.16.10.5 host 170.16.80.16
access-list 2 permit ip any any
interface Ethernet1/0
ip access-group 2 in
```
    
access-list 1 permit 192.168.2.2
interface Ethernet0/0
ip access-group 1 in

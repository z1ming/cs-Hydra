Specifications:

a) Traffic from the network 170.16.40.0 must not be allowed on the 170.16.50.0 network. All other traffic must be allowed on 170.16.50.0 as long as it originates from 170.16.0.0 (that is, outside traffic must not be allowed).

```
access-list 1 deny   170.16.40.0 0.0.0.255
access-list 1 permit 170.16.0.0  0.0.255.255
R2 # interface E0
ip access-group 1 out
```

b) Prevent all traffic from the workstation 170.16.10.5 from reaching the workstation 170.16.80.16. Traffic from all other hosts/ networks including traffic from outside should be allowed everywhere.

```
access-list 2 deny   170.16.10.5 0.0.0.0
access-list 2 permit any
R3 # interface E0
ip access-group 2 out
```

c) Workstations 170.16.50.75 and 170.16.50.7 should not be allowed HTTP access to the tower box 170.16.70.2. All other workstations can have HTTP access on the tower box. All other traffic including traffic from outside networks are allowed.

```
access-list 101 deny   tcp 170.16.50.75 0.0.0.0     170.16.70.2 0.0.0.0 eq 80
access-list 101 deny   tcp 170.16.50.7  0.0.0.0     170.16.70.2 0.0.0.0 eq 80
access-list 101 permit tcp 170.16.0.0   0.0.255.255 170.16.70.2 0.0.0.0 eq 80
access-list 101 permit ip  any          any
R3 # interface E1
ip access-group 101 out
```

d) 170.16.80.16 can telnet to 170.16.40.89. No one else from the network 170.16.80.0 can telnet to 170.16.40.89. Also permit all other traffic to 170.16.40.89, but only if they originate from 170.16.0.0 (that is, do not allow outside traffic).

```
access-list 102 permit tcp 170.16.80.16 0.0.0.0     170.16.40.0 0.0.0.255 eq 23
access-list 102 deny   tcp 170.16.80.0  0.0.0.255   170.16.40.0 0.0.0.255 eq 23
access-list 102 permit ip  170.16.0.0   0.0.255.255 170.16.40.0 0.0.0.255
R2 # interface E1
ip access-group 102 out
```


e) 170.16.10.5 can do only ftp access onto any host on the network 170.16.70.0. All other types of traffic from all other hosts are allowed, but only if they originate from 170.16.0.0 (that is, do not allow outside traffic).

```
access-list 103 permit tcp 170.16.10.5 0.0.0.0     170.16.70.0 0.0.0.255 range 20-21
access-list 103 permit ip  170.16.0.0  0.0.255.255 170.16.70.0 0.0.0.255
access-list 103 deny   ip  any         any
R3 # interface E1
ip access-group 103 out
```

f) Prevent traffic from the network 170.16.20.0 from flowing on the network 170.16.70.0 (that is, it must not flow on the network in either direction). All other traffic, including traffic from outside can.

```
access-list 3 deny   170.16.20.0 0.0.0.255
access-list 3 deny   170.16.70.0 0.0.0.255
access-list 3 permit any
R1 # interface E0
ip access-group 3 out
R3 # interface E1
ip access-group 3 out
```

g) Prevent traffic from the tower box 170.16.70.2 from going outside to the non-170.16.0.0 network. All other traffic can go out.

```
access-list 4 deny   170.16.70.2 0.0.0.0
access-list 4 permit any
R1 # interface S0
ip access-group 4 out
```

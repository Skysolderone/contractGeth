# contractGeth

solcjs --abi contract/myToken.sol --base-path . --include-path node_modules/ -o abi/

solcjs --bin contract/myToken.sol --base-path . --include-path node_modules/ -o bin/

abigen --abi=abi/contract_myToken_sol_WsToken.abi --bin=bin/contract_myToken_sol_WsToken.bin --pkg=gen --out=gen/m
yToken.go


testToken=0x9f0961F536c623F1432A242d13F49ca587D5104D //bscnet

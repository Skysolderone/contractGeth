# contractGeth

solcjs --abi contract/myToken.sol --base-path . --include-path node_modules/ -o abi/

solcjs --bin contract/myToken.sol --base-path . --include-path node_modules/ -o bin/

abigen --abi=abi/contract_myToken_sol_WsToken.abi --bin=bin/contract_myToken_sol_WsToken.bin --pkg=gen --out=gen/myToken.go

testToken=0xF1f5921C05620873163F35A25c533a1B3f3F47DE //bscnet

proxy
0x4c7659F3b5A2D0B2F35bE00C8F066899Bd85D4DB login
0xf638C3039D7b5BB55fee87F0D05C7d43fef3e9c4 proxy


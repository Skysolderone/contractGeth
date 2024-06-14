// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

contract LogicContract {
    uint public data;
    function SetData(uint _data) public {
        data = _data;
    }
    function GetData() public view returns (uint) {
        return data;
    }
}
contract ProxyContract {
    address public logic_contract;
    constructor(address _logic_contract) {
        logic_contract = _logic_contract;
    }
    fallback() external {
        address _impl = logic_contract;
        require(_impl != address(0));
        assembly {
            let ptr := mload(0x40)
            calldatacopy(ptr, 0, calldatasize())
            let result := delegatecall(gas(), _impl, ptr, calldatasize(), 0, 0)
            let size := returndatasize()
            returndatacopy(ptr, 0, size)
            switch result
            case 0 {
                revert(ptr, size)
            }
            default {
                return(ptr, size)
            }
        }
    }
}

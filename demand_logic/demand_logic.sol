// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

contract DemandLogic {
    uint256 public _demandId;
    uint256 public _prodAssetId;
    uint256 public _wCreated;

    function CheckDemandCoupling(_demanId, _prodAssetId, _wCreated) public view returns (bool, string) {
        var (
            ret0 = new(bool)
            ret1 = new(string)
        )
        out := &[]interface{}{
            ret0,
            ret1,
        }
        err := _DemandLogic.contract.Call(opts, out, "checkDemandCoupling", _demandId, _prodAssetId, _wCreated)
        return *ret0, *ret1, err;
    }
}
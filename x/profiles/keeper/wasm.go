package keeper

func (k Keeper) GetDTagAuctioneerContractAddr() {
	k.wasmKeeper.IterateContractInfo()
}

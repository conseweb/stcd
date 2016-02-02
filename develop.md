# 依赖关系
 stcd <-- go-flags 
 stcd <-- stcd/wire : test ok
 stcd <-- stcd/btcec : test ok
 stcd <-- stcd/btcjson : test ok
 stcd <-- stcd/addrmgr : test ok
 stcd/wire : test ok
 stcd/database <-- goleveldb : test ok
 stcd <-- stcd/blockchain <-- coinutil 
                          <-- stcd/database
                          <-- stcd/txscript
 stcd <-- stcd/peer <-- coinutil
 						<-- stcd/blockchain
 stcd <-- stcd/txscript <-- coinutil
 							 <-- stcd/database
 stcd/chaincfg <-- stcd/wire 			
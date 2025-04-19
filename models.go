package stakingwalletrpc

type Info struct {
	Balance           float64 `json:"balance,omitempty"`
	Blocks            uint64  `json:"blocks,omitempty"`
	Connections       uint8   `json:"connections,omitempty"`
	Difficulty        float64 `json:"difficulty,omitempty"`
	Errors            string  `json:"errors,omitempty"`
	Keypoololdest     uint64  `json:"keypoololdest,omitempty"`
	Keypoolsize       uint64  `json:"keypoolsize,omitempty"`
	Moneysupply       float64 `json:"moneysupply,omitempty"`
	Paytxfee          float64 `json:"paytxfee,omitempty"`
	Protocolversion   uint32  `json:"protocolversion,omitempty"`
	Proxy             string  `json:"proxy,omitempty"`
	Relayfee          float64 `json:"relayfee,omitempty"`
	Services          string  `json:"services,omitempty"`
	Shieldsuppy       float64 `json:"shieldsupply,omitempty"`
	Stakingstatus     string  `json:"staking status,omitempty"`
	Testnest          bool    `json:"testnest,omitempty"`
	Timeoffset        int8    `json:"timeoffset,omitempty"`
	Transparentsupply float64 `json:"transparentsupply,omitempty"`
	Version           int64   `json:"version,omitempty"`
	Walletversion     int64   `json:"walletversion,omitempty"`
}

type Transaction struct {
	Amount          float64               `json:"amount,omitempty"`
	Fee             float64               `json:"fee,omitempty"`
	Confirmations   uint64                `json:"confirmations"`
	Bcconfirmations uint64                `json:"bcconfirmations,omitempty"`
	Created         bool                  `json:"created,omitempty"`
	Blockhash       string                `json:"blockhash,omitempty"`
	Blockindex      uint8                 `json:"blockindex,omitempty"`
	Blocktime       uint64                `json:"blocktime,omitempty"`
	TxID            string                `json:"txid,omitempty"`
	Walletconflics  *[]WalletConflics     `json:"walletconflics,omitempty"`
	Time            uint64                `json:"time,omitempty"`
	Timereceived    uint64                `json:"timereceived,omitempty"`
	Details         *[]TransactionDetails `json:"details,omitempty"`
	Hex             string                `json:"hex,omitempty"`
}

type TransactionDetails struct {
	Address  string  `json:"address,omitempty"`
	Category string  `json:"category,omitempty"`
	Amount   float64 `json:"amount,omitempty"`
	Label    string  `json:"label,omitempty"`
	Vout     uint8   `json:"vout,omitempty"`
	Fee      float64 `json:"fee,omitempty"`
}

type WalletConflics struct {
}

type ValidateAddress struct {
	IsValid      bool    `json:"isvalid"`
	Address      *string `json:"address"`
	SciptPubKey  *string `json:"scriptPubKey"`
	IsMine       *bool   `json:"ismine"`
	IsWatchOnly  *bool   `json:"iswatchonly"`
	IsScript     *bool   `json:"isscript"`
	PubKey       *string `json:"pubkey"`
	IsCompressed *bool   `json:"iscompressed"`
	Label        *string `json:"label"`
}

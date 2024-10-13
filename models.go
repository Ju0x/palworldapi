package palworldapi

import "net/http"

type PalworldAPI struct {
	Host          string // Defaults to localhost:8212
	Username      string // Defaults to "admin"
	AdminPassword string

	client http.Client
}

// GET

type ServerInfo struct {
	Version     string `json:"version"`
	ServerName  string `json:"servername"`
	Description string `json:"description"`
}

type Player struct {
	Name        string  `json:"name"`
	AccountName string  `json:"accountName"`
	PlayerID    string  `json:"playerId"`
	UserID      string  `json:"userId"`
	IP          string  `json:"ip"`
	Ping        float64 `json:"ping"`
	LocationX   float64 `json:"location_x"`
	LocationY   float64 `json:"location_y"`
	Level       int     `json:"level"`
}

type PlayerList struct {
	Players []*Player `json:"players"`
}

type ServerSettings struct {
	Difficulty                          string `json:"Difficulty"`
	DayTimeSpeedRate                    int    `json:"DayTimeSpeedRate"`
	NightTimeSpeedRate                  int    `json:"NightTimeSpeedRate"`
	ExpRate                             int    `json:"ExpRate"`
	PalCaptureRate                      int    `json:"PalCaptureRate"`
	PalSpawnNumRate                     int    `json:"PalSpawnNumRate"`
	PalDamageRateAttack                 int    `json:"PalDamageRateAttack"`
	PalDamageRateDefense                int    `json:"PalDamageRateDefense"`
	PlayerDamageRateAttack              int    `json:"PlayerDamageRateAttack"`
	PlayerDamageRateDefense             int    `json:"PlayerDamageRateDefense"`
	PlayerStomachDecreaceRate           int    `json:"PlayerStomachDecreaceRate"`
	PlayerStaminaDecreaceRate           int    `json:"PlayerStaminaDecreaceRate"`
	PlayerAutoHPRegeneRate              int    `json:"PlayerAutoHPRegeneRate"`
	PlayerAutoHpRegeneRateInSleep       int    `json:"PlayerAutoHpRegeneRateInSleep"`
	PalStomachDecreaceRate              int    `json:"PalStomachDecreaceRate"`
	PalStaminaDecreaceRate              int    `json:"PalStaminaDecreaceRate"`
	PalAutoHPRegeneRate                 int    `json:"PalAutoHPRegeneRate"`
	PalAutoHpRegeneRateInSleep          int    `json:"PalAutoHpRegeneRateInSleep"`
	BuildObjectDamageRate               int    `json:"BuildObjectDamageRate"`
	BuildObjectDeteriorationDamageRate  int    `json:"BuildObjectDeteriorationDamageRate"`
	CollectionDropRate                  int    `json:"CollectionDropRate"`
	CollectionObjectHpRate              int    `json:"CollectionObjectHpRate"`
	CollectionObjectRespawnSpeedRate    int    `json:"CollectionObjectRespawnSpeedRate"`
	EnemyDropItemRate                   int    `json:"EnemyDropItemRate"`
	DeathPenalty                        string `json:"DeathPenalty"`
	EnablePlayerToPlayerDamage          bool   `json:"bEnablePlayerToPlayerDamage"`
	EnableFriendlyFire                  bool   `json:"bEnableFriendlyFire"`
	EnableInvaderEnemy                  bool   `json:"bEnableInvaderEnemy"`
	ActiveUNKO                          bool   `json:"bActiveUNKO"`
	EnableAimAssistPad                  bool   `json:"bEnableAimAssistPad"`
	EnableAimAssistKeyboard             bool   `json:"bEnableAimAssistKeyboard"`
	DropItemMaxNum                      int    `json:"DropItemMaxNum"`
	DropItemMaxNum_UNKO                 int    `json:"DropItemMaxNum_UNKO"`
	BaseCampMaxNum                      int    `json:"BaseCampMaxNum"`
	BaseCampWorkerMaxNum                int    `json:"BaseCampWorkerMaxNum"`
	DropItemAliveMaxHours               int    `json:"DropItemAliveMaxHours"`
	AutoResetGuildNoOnlinePlayers       bool   `json:"bAutoResetGuildNoOnlinePlayers"`
	AutoResetGuildTimeNoOnlinePlayers   int    `json:"AutoResetGuildTimeNoOnlinePlayers"`
	GuildPlayerMaxNum                   int    `json:"GuildPlayerMaxNum"`
	PalEggDefaultHatchingTime           int    `json:"PalEggDefaultHatchingTime"`
	WorkSpeedRate                       int    `json:"WorkSpeedRate"`
	IsMultiplay                         bool   `json:"bIsMultiplay"`
	IsPvP                               bool   `json:"bIsPvP"`
	CanPickupOtherGuildDeathPenaltyDrop bool   `json:"bCanPickupOtherGuildDeathPenaltyDrop"`
	EnableNonLoginPenalty               bool   `json:"bEnableNonLoginPenalty"`
	EnableFastTravel                    bool   `json:"bEnableFastTravel"`
	IsStartLocationSelectByMap          bool   `json:"bIsStartLocationSelectByMap"`
	ExistPlayerAfterLogout              bool   `json:"bExistPlayerAfterLogout"`
	EnableDefenseOtherGuildPlayer       bool   `json:"bEnableDefenseOtherGuildPlayer"`
	CoopPlayerMaxNum                    int    `json:"CoopPlayerMaxNum"`
	ServerPlayerMaxNum                  int    `json:"ServerPlayerMaxNum"`
	ServerName                          string `json:"ServerName"`
	ServerDescription                   string `json:"ServerDescription"`
	PublicPort                          int    `json:"PublicPort"`
	PublicIP                            string `json:"PublicIP"`
	RCONEnabled                         bool   `json:"RCONEnabled"`
	RCONPort                            int    `json:"RCONPort"`
	Region                              string `json:"Region"`
	UseAuth                             bool   `json:"bUseAuth"`
	BanListURL                          string `json:"BanListURL"`
	RESTAPIEnabled                      bool   `json:"RESTAPIEnabled"`
	RESTAPIPort                         int    `json:"RESTAPIPort"`
	BShowPlayerList                     bool   `json:"bShowPlayerList"`
	AllowConnectPlatform                string `json:"AllowConnectPlatform"`
	IsUseBackupSaveData                 bool   `json:"bIsUseBackupSaveData"`
	LogFormatType                       string `json:"LogFormatType"`
}

type ServerMetrics struct {
	FPS              int `json:"serverfps"`
	CurrentPlayerNum int `json:"currentplayernum"`
	MaxPlayerNum     int `json:"maxplayernum"`
	Uptime           int `json:"uptime"`
}

// POST

type Announce struct {
	Message string `json:"message"`
}

type KickPlayer struct {
	UserID  string `json:"userid"`
	Message string `json:"message"`
}

type BanPlayer struct {
	UserID  string `json:"userid"`
	Message string `json:"message"`
}

type UnbanPlayer struct {
	UserID string `json:"userid"`
}

type ServerShutdown struct {
	WaitTime int    `json:"waittime"` // Use custom datatype which converts time.Duration to the WaitTime
	Message  string `json:"string"`
}

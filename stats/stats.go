package stats

// BankAccount contains statistics about the player's money and spending.
type BankAccount struct {
	CurrentWealth          int64 `json:"Current_Wealth"`
	SpentOnShips           int64 `json:"Spent_On_Ships"`
	SpentOnOutfitting      int64 `json:"Spent_On_Outfitting"`
	SpentOnRepairs         int64 `json:"Spent_On_Repairs"`
	SpentOnFuel            int64 `json:"Spent_On_Fuel"`
	SpentOnAmmoConsumables int64 `json:"Spent_On_Ammo_Consumables"`
	InsuranceClaims        int64 `json:"Insurance_Claims"`
	SpentOnInsurance       int64 `json:"Spent_On_Insurance"`
	OwnedShipCount         int64 `json:"Owned_Ship_Count"`
}

// Combat contains statistics about combat and bounty.
type Combat struct {
	BountiesClaimed      int64 `json:"Bounties_Claimed"`
	BountyHuntingProfit  int64 `json:"Bounty_Hunting_Profit"`
	CombatBonds          int64 `json:"Combat_Bonds"`
	CombatBondProfits    int64 `json:"Combat_Bond_Profits"`
	Assassinations       int64 `json:"Assassinations"`
	AssassinationProfits int64 `json:"Assassination_Profits"`
	HighestSingleReward  int64 `json:"Highest_Single_Reward"`
	SkimmersKilled       int64 `json:"Skimmers_Killed"`
}

// Crime contains statistics about crime.
type Crime struct {
	Notoriety        int64 `json:"Notoriety"`
	Fines            int64 `json:"Fines"`
	TotalFines       int64 `json:"Total_Fines"`
	BountiesReceived int64 `json:"Bounties_Received"`
	TotalBounties    int64 `json:"Total_Bounties"`
	HighestBounty    int64 `json:"Highest_Bounty"`
}

// Smuggling contains statistics about smuggling.
type Smuggling struct {
	BlackMarketsTradedWith   int64   `json:"Black_Markets_Traded_With"`
	BlackMarketsProfits      int64   `json:"Black_Markets_Profits"`
	ResourcesSmuggled        int64   `json:"Resources_Smuggled"`
	AverageProfit            float64 `json:"Average_Profit"`
	HighestSingleTransaction int64   `json:"Highest_Single_Transaction"`
}

// Trading contains statistics about trading.
type Trading struct {
	MarketsTradedWith        int64   `json:"Markets_Traded_With"`
	MarketProfits            int64   `json:"Market_Profits"`
	ResourcesTraded          int64   `json:"Resources_Traded"`
	AverageProfit            float64 `json:"Average_Profit"`
	HighestSingleTransaction int64   `json:"Highest_Single_Transaction"`
}

// Mining contains statistics about mining.
type Mining struct {
	MiningProfits      int64 `json:"Mining_Profits"`
	QuantityMined      int64 `json:"Quantity_Mined"`
	MaterialsCollected int64 `json:"Materials_Collected"`
}

// Exploration contains statistics about exploration.
type Exploration struct {
	SystemsVisited            int64   `json:"Systems_Visited"`
	ExplorationProfits        int64   `json:"Exploration_Profits"`
	PlanetsScannedToLevel2    int64   `json:"Planets_Scanned_To_Level_2"`
	PlanetsScannedToLevel3    int64   `json:"Planets_Scanned_To_Level_3"`
	EfficientScans            int64   `json:"Efficient_Scans"`
	HighestPayout             int64   `json:"Highest_Payout"`
	TotalHyperspaceDistance   int64   `json:"Total_Hyperspace_Distance"`
	TotalHyperspaceJumps      int64   `json:"Total_Hyperspace_Jumps"`
	GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start"`
	TimePlayed                int64   `json:"Time_Played"`
}

// Passengers contains statistics about passenger missions.
type Passengers struct {
	PassengersMissionsAccepted  int64 `json:"Passengers_Missions_Accepted"`
	PassengersMissionsBulk      int64 `json:"Passengers_Missions_Bulk"`
	PassengersMissionsVIP       int64 `json:"Passengers_Missions_VIP"`
	PassengersMissionsDelivered int64 `json:"Passengers_Missions_Delivered"`
	PassengersMissionsEjected   int64 `json:"Passengers_Missions_Ejected"`
}

// SearchAndRescue contains statistics about search and rescue.
type SearchAndRescue struct {
	SearchRescueTraded int64 `json:"SearchRescue_Traded"`
	SearchRescueProfit int64 `json:"SearchRescue_Profit"`
	SearchRescueCount  int64 `json:"SearchRescue_Count"`
}

// Crafting contains statistics about crafting.
type Crafting struct {
	CountOfUsedEngineers  int64 `json:"Count_Of_Used_Engineers"`
	RecipesGenerated      int64 `json:"Recipes_Generated"`
	RecipesGeneratedRank1 int64 `json:"Recipes_Generated_Rank_1"`
	RecipesGeneratedRank2 int64 `json:"Recipes_Generated_Rank_2"`
	RecipesGeneratedRank3 int64 `json:"Recipes_Generated_Rank_3"`
	RecipesGeneratedRank4 int64 `json:"Recipes_Generated_Rank_4"`
	RecipesGeneratedRank5 int64 `json:"Recipes_Generated_Rank_5"`
}

// type Crew struct {
// }

// Multicrew contains statistics about multicrew.
type Multicrew struct {
	MulticrewTimeTotal        int64 `json:"Multicrew_Time_Total"`
	MulticrewGunnerTimeTotal  int64 `json:"Multicrew_Gunner_Time_Total"`
	MulticrewFighterTimeTotal int64 `json:"Multicrew_Fighter_Time_Total"`
	MulticrewCreditsTotal     int64 `json:"Multicrew_Credits_Total"`
	MulticrewFinesTotal       int64 `json:"Multicrew_Fines_Total"`
}

// MaterialTrader contains statistics about use of the Material Trader.
type MaterialTrader struct {
	TradesCompleted        int64 `json:"Trades_Completed"`
	MaterialsTraded        int64 `json:"Materials_Traded"`
	EncodedMaterialsTraded int64 `json:"Encoded_Materials_Traded"`
	Grade1MaterialsTraded  int64 `json:"Grade_1_Materials_Traded"`
	Grade2MaterialsTraded  int64 `json:"Grade_2_Materials_Traded"`
	Grade3MaterialsTraded  int64 `json:"Grade_3_Materials_Traded"`
	Grade4MaterialsTraded  int64 `json:"Grade_4_Materials_Traded"`
	Grade5MaterialsTraded  int64 `json:"Grade_5_Materials_Traded"`
}

package main

import (
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PacketPool map[uint32]func() packet.Packet

var Packets PacketPool

func init() {
	Packets = map[uint32]func() packet.Packet{
		packet.IDLogin:                      func() packet.Packet { return &packet.Login{} },
		packet.IDPlayStatus:                 func() packet.Packet { return &packet.PlayStatus{} },
		packet.IDServerToClientHandshake:    func() packet.Packet { return &packet.ServerToClientHandshake{} },
		packet.IDClientToServerHandshake:    func() packet.Packet { return &packet.ClientToServerHandshake{} },
		packet.IDDisconnect:                 func() packet.Packet { return &packet.Disconnect{} },
		packet.IDResourcePacksInfo:          func() packet.Packet { return &packet.ResourcePacksInfo{} },
		packet.IDResourcePackStack:          func() packet.Packet { return &packet.ResourcePackStack{} },
		packet.IDResourcePackClientResponse: func() packet.Packet { return &packet.ResourcePackClientResponse{} },
		packet.IDText:                       func() packet.Packet { return &packet.Text{} },
		packet.IDSetTime:                    func() packet.Packet { return &packet.SetTime{} },
		packet.IDStartGame:                  func() packet.Packet { return &packet.StartGame{} },
		packet.IDAddPlayer:                  func() packet.Packet { return &packet.AddPlayer{} },
		packet.IDAddActor:                   func() packet.Packet { return &packet.AddActor{} },
		packet.IDRemoveActor:                func() packet.Packet { return &packet.RemoveActor{} },
		packet.IDAddItemActor:               func() packet.Packet { return &packet.AddItemActor{} },
		// ---
		packet.IDTakeItemActor:     func() packet.Packet { return &packet.TakeItemActor{} },
		packet.IDMoveActorAbsolute: func() packet.Packet { return &packet.MoveActorAbsolute{} },
		packet.IDMovePlayer:        func() packet.Packet { return &packet.MovePlayer{} },
		packet.IDPassengerJump:     func() packet.Packet { return &packet.PassengerJump{} },
		packet.IDUpdateBlock:       func() packet.Packet { return &packet.UpdateBlock{} },
		packet.IDAddPainting:       func() packet.Packet { return &packet.AddPainting{} },
		packet.IDTickSync:          func() packet.Packet { return &packet.TickSync{} },
		// ---
		packet.IDLevelEvent:           func() packet.Packet { return &packet.LevelEvent{} },
		packet.IDBlockEvent:           func() packet.Packet { return &packet.BlockEvent{} },
		packet.IDActorEvent:           func() packet.Packet { return &packet.ActorEvent{} },
		packet.IDMobEffect:            func() packet.Packet { return &packet.MobEffect{} },
		packet.IDUpdateAttributes:     func() packet.Packet { return &packet.UpdateAttributes{} },
		packet.IDInventoryTransaction: func() packet.Packet { return &packet.InventoryTransaction{} },
		packet.IDMobEquipment:         func() packet.Packet { return &packet.MobEquipment{} },
		packet.IDMobArmourEquipment:   func() packet.Packet { return &packet.MobArmourEquipment{} },
		packet.IDInteract:             func() packet.Packet { return &packet.Interact{} },
		packet.IDBlockPickRequest:     func() packet.Packet { return &packet.BlockPickRequest{} },
		packet.IDActorPickRequest:     func() packet.Packet { return &packet.ActorPickRequest{} },
		packet.IDPlayerAction:         func() packet.Packet { return &packet.PlayerAction{} },
		// ---
		packet.IDHurtArmour:                  func() packet.Packet { return &packet.HurtArmour{} },
		packet.IDSetActorData:                func() packet.Packet { return &packet.SetActorData{} },
		packet.IDSetActorMotion:              func() packet.Packet { return &packet.SetActorMotion{} },
		packet.IDSetActorLink:                func() packet.Packet { return &packet.SetActorLink{} },
		packet.IDSetHealth:                   func() packet.Packet { return &packet.SetHealth{} },
		packet.IDSetSpawnPosition:            func() packet.Packet { return &packet.SetSpawnPosition{} },
		packet.IDAnimate:                     func() packet.Packet { return &packet.Animate{} },
		packet.IDRespawn:                     func() packet.Packet { return &packet.Respawn{} },
		packet.IDContainerOpen:               func() packet.Packet { return &packet.ContainerOpen{} },
		packet.IDContainerClose:              func() packet.Packet { return &packet.ContainerClose{} },
		packet.IDPlayerHotBar:                func() packet.Packet { return &packet.PlayerHotBar{} },
		packet.IDInventoryContent:            func() packet.Packet { return &packet.InventoryContent{} },
		packet.IDInventorySlot:               func() packet.Packet { return &packet.InventorySlot{} },
		packet.IDContainerSetData:            func() packet.Packet { return &packet.ContainerSetData{} },
		packet.IDCraftingData:                func() packet.Packet { return &packet.CraftingData{} },
		packet.IDCraftingEvent:               func() packet.Packet { return &packet.CraftingEvent{} },
		packet.IDGUIDataPickItem:             func() packet.Packet { return &packet.GUIDataPickItem{} },
		packet.IDAdventureSettings:           func() packet.Packet { return &packet.AdventureSettings{} },
		packet.IDBlockActorData:              func() packet.Packet { return &packet.BlockActorData{} },
		packet.IDPlayerInput:                 func() packet.Packet { return &packet.PlayerInput{} },
		packet.IDLevelChunk:                  func() packet.Packet { return &packet.LevelChunk{} },
		packet.IDSetCommandsEnabled:          func() packet.Packet { return &packet.SetCommandsEnabled{} },
		packet.IDSetDifficulty:               func() packet.Packet { return &packet.SetDifficulty{} },
		packet.IDChangeDimension:             func() packet.Packet { return &packet.ChangeDimension{} },
		packet.IDSetPlayerGameType:           func() packet.Packet { return &packet.SetPlayerGameType{} },
		packet.IDPlayerList:                  func() packet.Packet { return &packet.PlayerList{} },
		packet.IDSimpleEvent:                 func() packet.Packet { return &packet.SimpleEvent{} },
		packet.IDEvent:                       func() packet.Packet { return &packet.Event{} },
		packet.IDSpawnExperienceOrb:          func() packet.Packet { return &packet.SpawnExperienceOrb{} },
		packet.IDClientBoundMapItemData:      func() packet.Packet { return &packet.ClientBoundMapItemData{} },
		packet.IDMapInfoRequest:              func() packet.Packet { return &packet.MapInfoRequest{} },
		packet.IDRequestChunkRadius:          func() packet.Packet { return &packet.RequestChunkRadius{} },
		packet.IDChunkRadiusUpdated:          func() packet.Packet { return &packet.ChunkRadiusUpdated{} },
		packet.IDItemFrameDropItem:           func() packet.Packet { return &packet.ItemFrameDropItem{} },
		packet.IDGameRulesChanged:            func() packet.Packet { return &packet.GameRulesChanged{} },
		packet.IDCamera:                      func() packet.Packet { return &packet.Camera{} },
		packet.IDBossEvent:                   func() packet.Packet { return &packet.BossEvent{} },
		packet.IDShowCredits:                 func() packet.Packet { return &packet.ShowCredits{} },
		packet.IDAvailableCommands:           func() packet.Packet { return &packet.AvailableCommands{} },
		packet.IDCommandRequest:              func() packet.Packet { return &packet.CommandRequest{} },
		packet.IDCommandBlockUpdate:          func() packet.Packet { return &packet.CommandBlockUpdate{} },
		packet.IDCommandOutput:               func() packet.Packet { return &packet.CommandOutput{} },
		packet.IDUpdateTrade:                 func() packet.Packet { return &packet.UpdateTrade{} },
		packet.IDUpdateEquip:                 func() packet.Packet { return &packet.UpdateEquip{} },
		packet.IDResourcePackDataInfo:        func() packet.Packet { return &packet.ResourcePackDataInfo{} },
		packet.IDResourcePackChunkData:       func() packet.Packet { return &packet.ResourcePackChunkData{} },
		packet.IDResourcePackChunkRequest:    func() packet.Packet { return &packet.ResourcePackChunkRequest{} },
		packet.IDTransfer:                    func() packet.Packet { return &packet.Transfer{} },
		packet.IDPlaySound:                   func() packet.Packet { return &packet.PlaySound{} },
		packet.IDStopSound:                   func() packet.Packet { return &packet.StopSound{} },
		packet.IDSetTitle:                    func() packet.Packet { return &packet.SetTitle{} },
		packet.IDAddBehaviourTree:            func() packet.Packet { return &packet.AddBehaviourTree{} },
		packet.IDStructureBlockUpdate:        func() packet.Packet { return &packet.StructureBlockUpdate{} },
		packet.IDShowStoreOffer:              func() packet.Packet { return &packet.ShowStoreOffer{} },
		packet.IDPurchaseReceipt:             func() packet.Packet { return &packet.PurchaseReceipt{} },
		packet.IDPlayerSkin:                  func() packet.Packet { return &packet.PlayerSkin{} },
		packet.IDSubClientLogin:              func() packet.Packet { return &packet.SubClientLogin{} },
		packet.IDAutomationClientConnect:     func() packet.Packet { return &packet.AutomationClientConnect{} },
		packet.IDSetLastHurtBy:               func() packet.Packet { return &packet.SetLastHurtBy{} },
		packet.IDBookEdit:                    func() packet.Packet { return &packet.BookEdit{} },
		packet.IDNPCRequest:                  func() packet.Packet { return &packet.NPCRequest{} },
		packet.IDPhotoTransfer:               func() packet.Packet { return &packet.PhotoTransfer{} },
		packet.IDModalFormRequest:            func() packet.Packet { return &packet.ModalFormRequest{} },
		packet.IDModalFormResponse:           func() packet.Packet { return &packet.ModalFormResponse{} },
		packet.IDServerSettingsRequest:       func() packet.Packet { return &packet.ServerSettingsRequest{} },
		packet.IDServerSettingsResponse:      func() packet.Packet { return &packet.ServerSettingsResponse{} },
		packet.IDShowProfile:                 func() packet.Packet { return &packet.ShowProfile{} },
		packet.IDSetDefaultGameType:          func() packet.Packet { return &packet.SetDefaultGameType{} },
		packet.IDRemoveObjective:             func() packet.Packet { return &packet.RemoveObjective{} },
		packet.IDSetDisplayObjective:         func() packet.Packet { return &packet.SetDisplayObjective{} },
		packet.IDSetScore:                    func() packet.Packet { return &packet.SetScore{} },
		packet.IDLabTable:                    func() packet.Packet { return &packet.LabTable{} },
		packet.IDUpdateBlockSynced:           func() packet.Packet { return &packet.UpdateBlockSynced{} },
		packet.IDMoveActorDelta:              func() packet.Packet { return &packet.MoveActorDelta{} },
		packet.IDSetScoreboardIdentity:       func() packet.Packet { return &packet.SetScoreboardIdentity{} },
		packet.IDSetLocalPlayerAsInitialised: func() packet.Packet { return &packet.SetLocalPlayerAsInitialised{} },
		packet.IDUpdateSoftEnum:              func() packet.Packet { return &packet.UpdateSoftEnum{} },
		packet.IDNetworkStackLatency:         func() packet.Packet { return &packet.NetworkStackLatency{} },
		// ---
		packet.IDScriptCustomEvent:           func() packet.Packet { return &packet.ScriptCustomEvent{} },
		packet.IDSpawnParticleEffect:         func() packet.Packet { return &packet.SpawnParticleEffect{} },
		packet.IDAvailableActorIdentifiers:   func() packet.Packet { return &packet.AvailableActorIdentifiers{} },
		packet.IDNetworkChunkPublisherUpdate: func() packet.Packet { return &packet.NetworkChunkPublisherUpdate{} },
		packet.IDBiomeDefinitionList:         func() packet.Packet { return &packet.BiomeDefinitionList{} },
		packet.IDLevelSoundEvent:             func() packet.Packet { return &packet.LevelSoundEvent{} },
		packet.IDLevelEventGeneric:           func() packet.Packet { return &packet.LevelEventGeneric{} },
		packet.IDLecternUpdate:               func() packet.Packet { return &packet.LecternUpdate{} },
		// ---
		packet.IDAddEntity:                     func() packet.Packet { return &packet.AddEntity{} },
		packet.IDRemoveEntity:                  func() packet.Packet { return &packet.RemoveEntity{} },
		packet.IDClientCacheStatus:             func() packet.Packet { return &packet.ClientCacheStatus{} },
		packet.IDOnScreenTextureAnimation:      func() packet.Packet { return &packet.OnScreenTextureAnimation{} },
		packet.IDMapCreateLockedCopy:           func() packet.Packet { return &packet.MapCreateLockedCopy{} },
		packet.IDStructureTemplateDataRequest:  func() packet.Packet { return &packet.StructureTemplateDataResponse{} },
		packet.IDStructureTemplateDataResponse: func() packet.Packet { return &packet.StructureTemplateDataResponse{} },
		// ---
		packet.IDClientCacheBlobStatus:             func() packet.Packet { return &packet.ClientCacheBlobStatus{} },
		packet.IDClientCacheMissResponse:           func() packet.Packet { return &packet.ClientCacheMissResponse{} },
		packet.IDEducationSettings:                 func() packet.Packet { return &packet.EducationSettings{} },
		packet.IDEmote:                             func() packet.Packet { return &packet.Emote{} },
		packet.IDMultiPlayerSettings:               func() packet.Packet { return &packet.MultiPlayerSettings{} },
		packet.IDSettingsCommand:                   func() packet.Packet { return &packet.SettingsCommand{} },
		packet.IDAnvilDamage:                       func() packet.Packet { return &packet.AnvilDamage{} },
		packet.IDCompletedUsingItem:                func() packet.Packet { return &packet.CompletedUsingItem{} },
		packet.IDNetworkSettings:                   func() packet.Packet { return &packet.NetworkSettings{} },
		packet.IDPlayerAuthInput:                   func() packet.Packet { return &packet.PlayerAuthInput{} },
		packet.IDCreativeContent:                   func() packet.Packet { return &packet.CreativeContent{} },
		packet.IDPlayerEnchantOptions:              func() packet.Packet { return &packet.PlayerEnchantOptions{} },
		packet.IDItemStackRequest:                  func() packet.Packet { return &packet.ItemStackRequest{} },
		packet.IDItemStackResponse:                 func() packet.Packet { return &packet.ItemStackResponse{} },
		packet.IDPlayerArmourDamage:                func() packet.Packet { return &packet.PlayerArmourDamage{} },
		packet.IDCodeBuilder:                       func() packet.Packet { return &packet.CodeBuilder{} },
		packet.IDUpdatePlayerGameType:              func() packet.Packet { return &packet.UpdatePlayerGameType{} },
		packet.IDEmoteList:                         func() packet.Packet { return &packet.EmoteList{} },
		packet.IDPositionTrackingDBServerBroadcast: func() packet.Packet { return &packet.PositionTrackingDBServerBroadcast{} },
		packet.IDPositionTrackingDBClientRequest:   func() packet.Packet { return &packet.PositionTrackingDBClientRequest{} },
		packet.IDDebugInfo:                         func() packet.Packet { return &packet.DebugInfo{} },
		packet.IDPacketViolationWarning:            func() packet.Packet { return &packet.PacketViolationWarning{} },
		packet.IDMotionPredictionHints:             func() packet.Packet { return &packet.MotionPredictionHints{} },
		packet.IDAnimateEntity:                     func() packet.Packet { return &packet.AnimateEntity{} },
		packet.IDCameraShake:                       func() packet.Packet { return &packet.CameraShake{} },
		packet.IDPlayerFog:                         func() packet.Packet { return &packet.PlayerFog{} },
		packet.IDCorrectPlayerMovePrediction:       func() packet.Packet { return &packet.CorrectPlayerMovePrediction{} },
		packet.IDItemComponent:                     func() packet.Packet { return &packet.ItemComponent{} },
		packet.IDFilterText:                        func() packet.Packet { return &packet.FilterText{} },
		packet.IDClientBoundDebugRenderer:          func() packet.Packet { return &packet.ClientBoundDebugRenderer{} },
		packet.IDSyncActorProperty:                 func() packet.Packet { return &packet.SyncActorProperty{} },
		packet.IDAddVolumeEntity:                   func() packet.Packet { return &packet.AddVolumeEntity{} },
		packet.IDRemoveVolumeEntity:                func() packet.Packet { return &packet.RemoveVolumeEntity{} },
		packet.IDSimulationType:                    func() packet.Packet { return &packet.SimulationType{} },
		packet.IDNPCDialogue:                       func() packet.Packet { return &packet.NPCDialogue{} },
		packet.IDEducationResourceURI:              func() packet.Packet { return &packet.EducationResourceURI{} },
		packet.IDCreatePhoto:                       func() packet.Packet { return &packet.CreatePhoto{} },
		packet.IDUpdateSubChunkBlocks:              func() packet.Packet { return &packet.UpdateSubChunkBlocks{} },
		packet.IDPhotoInfoRequest:                  func() packet.Packet { return &packet.PhotoInfoRequest{} },
		packet.IDSubChunk:                          func() packet.Packet { return &packet.SubChunk{} },
		packet.IDSubChunkRequest:                   func() packet.Packet { return &packet.SubChunkRequest{} },
	}
}
package item

type Slot uint8
type ArmorSlot Slot
type EquipSlot Slot
type AugmentSlot Slot

const (
	SlotNone Slot = 0xFF

	EquipSlotMisc1 EquipSlot = 0x01
	EquipSlotMisc2 EquipSlot = 0x02
	EquipSlotMisc3 EquipSlot = 0x03
	EquipSlotMisc4 EquipSlot = 0x04
	EquipSlotMisc5 EquipSlot = 0x05
	EquipSlotMisc6 EquipSlot = 0x06
	EquipRightHand EquipSlot = 0x07
	EquipCore      EquipSlot = 0x08
	EquipLeftHand  EquipSlot = 0x09
	EquipCoreSlot1 EquipSlot = 0x0A
	EquipCoreSlot2 EquipSlot = 0x0B
	EquipCoreSlot3 EquipSlot = 0x0C
	EquipCoreSlot4 EquipSlot = 0x0D
	EquipCoreSlot5 EquipSlot = 0x0E
	EquipCoreSlot6 EquipSlot = 0x0F

	ArmorSlotNeck          ArmorSlot = 0x11
	ArmorSlotHead          ArmorSlot = 0x12
	ArmorSlotShoulders     ArmorSlot = 0x13
	ArmorSlotRightShoulder ArmorSlot = 0x14
	ArmorSlotChest         ArmorSlot = 0x15
	ArmorSlotLeftShoulder  ArmorSlot = 0x16
	ArmorSlotRightArm      ArmorSlot = 0x17
	ArmorSlotHands         ArmorSlot = 0x18
	ArmorSlotLeftArm       ArmorSlot = 0x19
	ArmorSlotRightHand     ArmorSlot = 0x1A
	ArmorSlotLegs          ArmorSlot = 0x1B
	ArmorSlotLeftHand      ArmorSlot = 0x1C
	ArmorSlotAura1         ArmorSlot = 0x1D
	ArmorSlotFeet          ArmorSlot = 0x1E
	ArmorSlotAura2         ArmorSlot = 0x1F

	AugmentSlot1  AugmentSlot = 0x21
	AugmentSlot2  AugmentSlot = 0x22
	AugmentSlot3  AugmentSlot = 0x23
	AugmentSlot4  AugmentSlot = 0x24
	AugmentSlot5  AugmentSlot = 0x25
	AugmentSlot6  AugmentSlot = 0x26
	AugmentSlot7  AugmentSlot = 0x27
	AugmentSlot8  AugmentSlot = 0x28
	AugmentSlot9  AugmentSlot = 0x29
	AugmentSlot10 AugmentSlot = 0x2A
	AugmentSlot11 AugmentSlot = 0x2B
	AugmentSlot12 AugmentSlot = 0x2C
	AugmentSlot13 AugmentSlot = 0x2D
	AugmentSlot14 AugmentSlot = 0x2E
	AugmentSlot15 AugmentSlot = 0x2F
)

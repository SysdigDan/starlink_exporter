// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spacex/api/device/transceiver.proto

package device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TransceiverModulatorState int32

const (
	TransceiverModulatorState_MODSTATE_UNKNOWN  TransceiverModulatorState = 0
	TransceiverModulatorState_MODSTATE_ENABLED  TransceiverModulatorState = 1
	TransceiverModulatorState_MODSTATE_DISABLED TransceiverModulatorState = 2
)

var TransceiverModulatorState_name = map[int32]string{
	0: "MODSTATE_UNKNOWN",
	1: "MODSTATE_ENABLED",
	2: "MODSTATE_DISABLED",
}

var TransceiverModulatorState_value = map[string]int32{
	"MODSTATE_UNKNOWN":  0,
	"MODSTATE_ENABLED":  1,
	"MODSTATE_DISABLED": 2,
}

func (x TransceiverModulatorState) String() string {
	return proto.EnumName(TransceiverModulatorState_name, int32(x))
}

func (TransceiverModulatorState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{0}
}

type TransceiverTxRxState int32

const (
	TransceiverTxRxState_TXRX_UNKNOWN  TransceiverTxRxState = 0
	TransceiverTxRxState_TXRX_ENABLED  TransceiverTxRxState = 1
	TransceiverTxRxState_TXRX_DISABLED TransceiverTxRxState = 2
)

var TransceiverTxRxState_name = map[int32]string{
	0: "TXRX_UNKNOWN",
	1: "TXRX_ENABLED",
	2: "TXRX_DISABLED",
}

var TransceiverTxRxState_value = map[string]int32{
	"TXRX_UNKNOWN":  0,
	"TXRX_ENABLED":  1,
	"TXRX_DISABLED": 2,
}

func (x TransceiverTxRxState) String() string {
	return proto.EnumName(TransceiverTxRxState_name, int32(x))
}

func (TransceiverTxRxState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{1}
}

type TransceiverTransmitBlankingState int32

const (
	TransceiverTransmitBlankingState_TB_UNKNOWN  TransceiverTransmitBlankingState = 0
	TransceiverTransmitBlankingState_TB_ENABLED  TransceiverTransmitBlankingState = 1
	TransceiverTransmitBlankingState_TB_DISABLED TransceiverTransmitBlankingState = 2
)

var TransceiverTransmitBlankingState_name = map[int32]string{
	0: "TB_UNKNOWN",
	1: "TB_ENABLED",
	2: "TB_DISABLED",
}

var TransceiverTransmitBlankingState_value = map[string]int32{
	"TB_UNKNOWN":  0,
	"TB_ENABLED":  1,
	"TB_DISABLED": 2,
}

func (x TransceiverTransmitBlankingState) String() string {
	return proto.EnumName(TransceiverTransmitBlankingState_name, int32(x))
}

func (TransceiverTransmitBlankingState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{2}
}

type TransceiverIFLoopbackTestRequest struct {
	EnableIfLoopback     bool     `protobuf:"varint,1,opt,name=enable_if_loopback,json=enableIfLoopback,proto3" json:"enable_if_loopback,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransceiverIFLoopbackTestRequest) Reset()         { *m = TransceiverIFLoopbackTestRequest{} }
func (m *TransceiverIFLoopbackTestRequest) String() string { return proto.CompactTextString(m) }
func (*TransceiverIFLoopbackTestRequest) ProtoMessage()    {}
func (*TransceiverIFLoopbackTestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{0}
}

func (m *TransceiverIFLoopbackTestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverIFLoopbackTestRequest.Unmarshal(m, b)
}
func (m *TransceiverIFLoopbackTestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverIFLoopbackTestRequest.Marshal(b, m, deterministic)
}
func (m *TransceiverIFLoopbackTestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverIFLoopbackTestRequest.Merge(m, src)
}
func (m *TransceiverIFLoopbackTestRequest) XXX_Size() int {
	return xxx_messageInfo_TransceiverIFLoopbackTestRequest.Size(m)
}
func (m *TransceiverIFLoopbackTestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverIFLoopbackTestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverIFLoopbackTestRequest proto.InternalMessageInfo

func (m *TransceiverIFLoopbackTestRequest) GetEnableIfLoopback() bool {
	if m != nil {
		return m.EnableIfLoopback
	}
	return false
}

type TransceiverIFLoopbackTestResponse struct {
	BerLoopbackTest      float32  `protobuf:"fixed32,1,opt,name=ber_loopback_test,json=berLoopbackTest,proto3" json:"ber_loopback_test,omitempty"`
	SnrLoopbackTest      float32  `protobuf:"fixed32,2,opt,name=snr_loopback_test,json=snrLoopbackTest,proto3" json:"snr_loopback_test,omitempty"`
	RssiLoopbackTest     float32  `protobuf:"fixed32,3,opt,name=rssi_loopback_test,json=rssiLoopbackTest,proto3" json:"rssi_loopback_test,omitempty"`
	PllLock              bool     `protobuf:"varint,4,opt,name=pll_lock,json=pllLock,proto3" json:"pll_lock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransceiverIFLoopbackTestResponse) Reset()         { *m = TransceiverIFLoopbackTestResponse{} }
func (m *TransceiverIFLoopbackTestResponse) String() string { return proto.CompactTextString(m) }
func (*TransceiverIFLoopbackTestResponse) ProtoMessage()    {}
func (*TransceiverIFLoopbackTestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{1}
}

func (m *TransceiverIFLoopbackTestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverIFLoopbackTestResponse.Unmarshal(m, b)
}
func (m *TransceiverIFLoopbackTestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverIFLoopbackTestResponse.Marshal(b, m, deterministic)
}
func (m *TransceiverIFLoopbackTestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverIFLoopbackTestResponse.Merge(m, src)
}
func (m *TransceiverIFLoopbackTestResponse) XXX_Size() int {
	return xxx_messageInfo_TransceiverIFLoopbackTestResponse.Size(m)
}
func (m *TransceiverIFLoopbackTestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverIFLoopbackTestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverIFLoopbackTestResponse proto.InternalMessageInfo

func (m *TransceiverIFLoopbackTestResponse) GetBerLoopbackTest() float32 {
	if m != nil {
		return m.BerLoopbackTest
	}
	return 0
}

func (m *TransceiverIFLoopbackTestResponse) GetSnrLoopbackTest() float32 {
	if m != nil {
		return m.SnrLoopbackTest
	}
	return 0
}

func (m *TransceiverIFLoopbackTestResponse) GetRssiLoopbackTest() float32 {
	if m != nil {
		return m.RssiLoopbackTest
	}
	return 0
}

func (m *TransceiverIFLoopbackTestResponse) GetPllLock() bool {
	if m != nil {
		return m.PllLock
	}
	return false
}

type TransceiverGetStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransceiverGetStatusRequest) Reset()         { *m = TransceiverGetStatusRequest{} }
func (m *TransceiverGetStatusRequest) String() string { return proto.CompactTextString(m) }
func (*TransceiverGetStatusRequest) ProtoMessage()    {}
func (*TransceiverGetStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{2}
}

func (m *TransceiverGetStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverGetStatusRequest.Unmarshal(m, b)
}
func (m *TransceiverGetStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverGetStatusRequest.Marshal(b, m, deterministic)
}
func (m *TransceiverGetStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverGetStatusRequest.Merge(m, src)
}
func (m *TransceiverGetStatusRequest) XXX_Size() int {
	return xxx_messageInfo_TransceiverGetStatusRequest.Size(m)
}
func (m *TransceiverGetStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverGetStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverGetStatusRequest proto.InternalMessageInfo

type TransceiverGetStatusResponse struct {
	ModState              TransceiverModulatorState        `protobuf:"varint,1,opt,name=mod_state,json=modState,proto3,enum=SpaceX.API.Device.TransceiverModulatorState" json:"mod_state,omitempty"`
	DemodState            TransceiverModulatorState        `protobuf:"varint,2,opt,name=demod_state,json=demodState,proto3,enum=SpaceX.API.Device.TransceiverModulatorState" json:"demod_state,omitempty"`
	TxState               TransceiverTxRxState             `protobuf:"varint,3,opt,name=tx_state,json=txState,proto3,enum=SpaceX.API.Device.TransceiverTxRxState" json:"tx_state,omitempty"`
	RxState               TransceiverTxRxState             `protobuf:"varint,4,opt,name=rx_state,json=rxState,proto3,enum=SpaceX.API.Device.TransceiverTxRxState" json:"rx_state,omitempty"`
	State                 DishState                        `protobuf:"varint,1006,opt,name=state,proto3,enum=SpaceX.API.Device.DishState" json:"state,omitempty"`
	Faults                *TransceiverFaults               `protobuf:"bytes,1007,opt,name=faults,proto3" json:"faults,omitempty"`
	TransmitBlankingState TransceiverTransmitBlankingState `protobuf:"varint,1008,opt,name=transmit_blanking_state,json=transmitBlankingState,proto3,enum=SpaceX.API.Device.TransceiverTransmitBlankingState" json:"transmit_blanking_state,omitempty"`
	ModemAsicTemp         float32                          `protobuf:"fixed32,1009,opt,name=modem_asic_temp,json=modemAsicTemp,proto3" json:"modem_asic_temp,omitempty"`
	TxIfTemp              float32                          `protobuf:"fixed32,1010,opt,name=tx_if_temp,json=txIfTemp,proto3" json:"tx_if_temp,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                         `json:"-"`
	XXX_unrecognized      []byte                           `json:"-"`
	XXX_sizecache         int32                            `json:"-"`
}

func (m *TransceiverGetStatusResponse) Reset()         { *m = TransceiverGetStatusResponse{} }
func (m *TransceiverGetStatusResponse) String() string { return proto.CompactTextString(m) }
func (*TransceiverGetStatusResponse) ProtoMessage()    {}
func (*TransceiverGetStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{3}
}

func (m *TransceiverGetStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverGetStatusResponse.Unmarshal(m, b)
}
func (m *TransceiverGetStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverGetStatusResponse.Marshal(b, m, deterministic)
}
func (m *TransceiverGetStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverGetStatusResponse.Merge(m, src)
}
func (m *TransceiverGetStatusResponse) XXX_Size() int {
	return xxx_messageInfo_TransceiverGetStatusResponse.Size(m)
}
func (m *TransceiverGetStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverGetStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverGetStatusResponse proto.InternalMessageInfo

func (m *TransceiverGetStatusResponse) GetModState() TransceiverModulatorState {
	if m != nil {
		return m.ModState
	}
	return TransceiverModulatorState_MODSTATE_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetDemodState() TransceiverModulatorState {
	if m != nil {
		return m.DemodState
	}
	return TransceiverModulatorState_MODSTATE_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetTxState() TransceiverTxRxState {
	if m != nil {
		return m.TxState
	}
	return TransceiverTxRxState_TXRX_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetRxState() TransceiverTxRxState {
	if m != nil {
		return m.RxState
	}
	return TransceiverTxRxState_TXRX_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetState() DishState {
	if m != nil {
		return m.State
	}
	return DishState_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetFaults() *TransceiverFaults {
	if m != nil {
		return m.Faults
	}
	return nil
}

func (m *TransceiverGetStatusResponse) GetTransmitBlankingState() TransceiverTransmitBlankingState {
	if m != nil {
		return m.TransmitBlankingState
	}
	return TransceiverTransmitBlankingState_TB_UNKNOWN
}

func (m *TransceiverGetStatusResponse) GetModemAsicTemp() float32 {
	if m != nil {
		return m.ModemAsicTemp
	}
	return 0
}

func (m *TransceiverGetStatusResponse) GetTxIfTemp() float32 {
	if m != nil {
		return m.TxIfTemp
	}
	return 0
}

type TransceiverFaults struct {
	OverTempModemAsicFault bool     `protobuf:"varint,1,opt,name=over_temp_modem_asic_fault,json=overTempModemAsicFault,proto3" json:"over_temp_modem_asic_fault,omitempty"`
	OverTempPcbaFault      bool     `protobuf:"varint,2,opt,name=over_temp_pcba_fault,json=overTempPcbaFault,proto3" json:"over_temp_pcba_fault,omitempty"`
	DcVoltageFault         bool     `protobuf:"varint,3,opt,name=dc_voltage_fault,json=dcVoltageFault,proto3" json:"dc_voltage_fault,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *TransceiverFaults) Reset()         { *m = TransceiverFaults{} }
func (m *TransceiverFaults) String() string { return proto.CompactTextString(m) }
func (*TransceiverFaults) ProtoMessage()    {}
func (*TransceiverFaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{4}
}

func (m *TransceiverFaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverFaults.Unmarshal(m, b)
}
func (m *TransceiverFaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverFaults.Marshal(b, m, deterministic)
}
func (m *TransceiverFaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverFaults.Merge(m, src)
}
func (m *TransceiverFaults) XXX_Size() int {
	return xxx_messageInfo_TransceiverFaults.Size(m)
}
func (m *TransceiverFaults) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverFaults.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverFaults proto.InternalMessageInfo

func (m *TransceiverFaults) GetOverTempModemAsicFault() bool {
	if m != nil {
		return m.OverTempModemAsicFault
	}
	return false
}

func (m *TransceiverFaults) GetOverTempPcbaFault() bool {
	if m != nil {
		return m.OverTempPcbaFault
	}
	return false
}

func (m *TransceiverFaults) GetDcVoltageFault() bool {
	if m != nil {
		return m.DcVoltageFault
	}
	return false
}

type TransceiverGetTelemetryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransceiverGetTelemetryRequest) Reset()         { *m = TransceiverGetTelemetryRequest{} }
func (m *TransceiverGetTelemetryRequest) String() string { return proto.CompactTextString(m) }
func (*TransceiverGetTelemetryRequest) ProtoMessage()    {}
func (*TransceiverGetTelemetryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{5}
}

func (m *TransceiverGetTelemetryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverGetTelemetryRequest.Unmarshal(m, b)
}
func (m *TransceiverGetTelemetryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverGetTelemetryRequest.Marshal(b, m, deterministic)
}
func (m *TransceiverGetTelemetryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverGetTelemetryRequest.Merge(m, src)
}
func (m *TransceiverGetTelemetryRequest) XXX_Size() int {
	return xxx_messageInfo_TransceiverGetTelemetryRequest.Size(m)
}
func (m *TransceiverGetTelemetryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverGetTelemetryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverGetTelemetryRequest proto.InternalMessageInfo

type TransceiverGetTelemetryResponse struct {
	AntennaPointingMode                uint32   `protobuf:"varint,1001,opt,name=antenna_pointing_mode,json=antennaPointingMode,proto3" json:"antenna_pointing_mode,omitempty"`
	AntennaPitch                       float32  `protobuf:"fixed32,1002,opt,name=antenna_pitch,json=antennaPitch,proto3" json:"antenna_pitch,omitempty"`
	AntennaRoll                        float32  `protobuf:"fixed32,1003,opt,name=antenna_roll,json=antennaRoll,proto3" json:"antenna_roll,omitempty"`
	AntennaRxTheta                     float32  `protobuf:"fixed32,1004,opt,name=antenna_rx_theta,json=antennaRxTheta,proto3" json:"antenna_rx_theta,omitempty"`
	AntennaTrueHeading                 float32  `protobuf:"fixed32,1005,opt,name=antenna_true_heading,json=antennaTrueHeading,proto3" json:"antenna_true_heading,omitempty"`
	RxChannel                          uint32   `protobuf:"varint,1006,opt,name=rx_channel,json=rxChannel,proto3" json:"rx_channel,omitempty"`
	CurrentCellId                      uint32   `protobuf:"varint,1007,opt,name=current_cell_id,json=currentCellId,proto3" json:"current_cell_id,omitempty"`
	SecondsUntilSlotEnd                float32  `protobuf:"fixed32,1008,opt,name=seconds_until_slot_end,json=secondsUntilSlotEnd,proto3" json:"seconds_until_slot_end,omitempty"`
	WbRssiPeakMagDb                    float32  `protobuf:"fixed32,1009,opt,name=wb_rssi_peak_mag_db,json=wbRssiPeakMagDb,proto3" json:"wb_rssi_peak_mag_db,omitempty"`
	PopPingDropRate                    float32  `protobuf:"fixed32,1010,opt,name=pop_ping_drop_rate,json=popPingDropRate,proto3" json:"pop_ping_drop_rate,omitempty"`
	SnrDb                              float32  `protobuf:"fixed32,1011,opt,name=snr_db,json=snrDb,proto3" json:"snr_db,omitempty"`
	L1SnrAvgDb                         float32  `protobuf:"fixed32,1012,opt,name=l1_snr_avg_db,json=l1SnrAvgDb,proto3" json:"l1_snr_avg_db,omitempty"`
	L1SnrMinDb                         float32  `protobuf:"fixed32,1013,opt,name=l1_snr_min_db,json=l1SnrMinDb,proto3" json:"l1_snr_min_db,omitempty"`
	L1SnrMaxDb                         float32  `protobuf:"fixed32,1014,opt,name=l1_snr_max_db,json=l1SnrMaxDb,proto3" json:"l1_snr_max_db,omitempty"`
	LmacSatelliteId                    uint32   `protobuf:"varint,1015,opt,name=lmac_satellite_id,json=lmacSatelliteId,proto3" json:"lmac_satellite_id,omitempty"`
	TargetSatelliteId                  uint32   `protobuf:"varint,1016,opt,name=target_satellite_id,json=targetSatelliteId,proto3" json:"target_satellite_id,omitempty"`
	GrantMcs                           uint32   `protobuf:"varint,1017,opt,name=grant_mcs,json=grantMcs,proto3" json:"grant_mcs,omitempty"`
	GrantSymbolsAvg                    float32  `protobuf:"fixed32,1018,opt,name=grant_symbols_avg,json=grantSymbolsAvg,proto3" json:"grant_symbols_avg,omitempty"`
	DedGrant                           uint32   `protobuf:"varint,1019,opt,name=ded_grant,json=dedGrant,proto3" json:"ded_grant,omitempty"`
	MobilityProactiveSlotChange        uint32   `protobuf:"varint,1020,opt,name=mobility_proactive_slot_change,json=mobilityProactiveSlotChange,proto3" json:"mobility_proactive_slot_change,omitempty"`
	MobilityReactiveSlotChange         uint32   `protobuf:"varint,1021,opt,name=mobility_reactive_slot_change,json=mobilityReactiveSlotChange,proto3" json:"mobility_reactive_slot_change,omitempty"`
	RfpTotalSynFailed                  uint32   `protobuf:"varint,1022,opt,name=rfp_total_syn_failed,json=rfpTotalSynFailed,proto3" json:"rfp_total_syn_failed,omitempty"`
	NumOutOfSeq                        uint32   `protobuf:"varint,1023,opt,name=num_out_of_seq,json=numOutOfSeq,proto3" json:"num_out_of_seq,omitempty"`
	NumUlmapDrop                       uint32   `protobuf:"varint,1024,opt,name=num_ulmap_drop,json=numUlmapDrop,proto3" json:"num_ulmap_drop,omitempty"`
	CurrentSecondsOfSchedule           float32  `protobuf:"fixed32,1025,opt,name=current_seconds_of_schedule,json=currentSecondsOfSchedule,proto3" json:"current_seconds_of_schedule,omitempty"`
	SendLabelSwitchToGroundFailedCalls uint32   `protobuf:"varint,1026,opt,name=send_label_switch_to_ground_failed_calls,json=sendLabelSwitchToGroundFailedCalls,proto3" json:"send_label_switch_to_ground_failed_calls,omitempty"`
	EmaVelocityX                       float64  `protobuf:"fixed64,1027,opt,name=ema_velocity_x,json=emaVelocityX,proto3" json:"ema_velocity_x,omitempty"`
	EmaVelocityY                       float64  `protobuf:"fixed64,1028,opt,name=ema_velocity_y,json=emaVelocityY,proto3" json:"ema_velocity_y,omitempty"`
	EmaVelocityZ                       float64  `protobuf:"fixed64,1029,opt,name=ema_velocity_z,json=emaVelocityZ,proto3" json:"ema_velocity_z,omitempty"`
	CeRssiDb                           float32  `protobuf:"fixed32,1030,opt,name=ce_rssi_db,json=ceRssiDb,proto3" json:"ce_rssi_db,omitempty"`
	XXX_NoUnkeyedLiteral               struct{} `json:"-"`
	XXX_unrecognized                   []byte   `json:"-"`
	XXX_sizecache                      int32    `json:"-"`
}

func (m *TransceiverGetTelemetryResponse) Reset()         { *m = TransceiverGetTelemetryResponse{} }
func (m *TransceiverGetTelemetryResponse) String() string { return proto.CompactTextString(m) }
func (*TransceiverGetTelemetryResponse) ProtoMessage()    {}
func (*TransceiverGetTelemetryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_528fde07cc1e136a, []int{6}
}

func (m *TransceiverGetTelemetryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransceiverGetTelemetryResponse.Unmarshal(m, b)
}
func (m *TransceiverGetTelemetryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransceiverGetTelemetryResponse.Marshal(b, m, deterministic)
}
func (m *TransceiverGetTelemetryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransceiverGetTelemetryResponse.Merge(m, src)
}
func (m *TransceiverGetTelemetryResponse) XXX_Size() int {
	return xxx_messageInfo_TransceiverGetTelemetryResponse.Size(m)
}
func (m *TransceiverGetTelemetryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransceiverGetTelemetryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransceiverGetTelemetryResponse proto.InternalMessageInfo

func (m *TransceiverGetTelemetryResponse) GetAntennaPointingMode() uint32 {
	if m != nil {
		return m.AntennaPointingMode
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetAntennaPitch() float32 {
	if m != nil {
		return m.AntennaPitch
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetAntennaRoll() float32 {
	if m != nil {
		return m.AntennaRoll
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetAntennaRxTheta() float32 {
	if m != nil {
		return m.AntennaRxTheta
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetAntennaTrueHeading() float32 {
	if m != nil {
		return m.AntennaTrueHeading
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetRxChannel() uint32 {
	if m != nil {
		return m.RxChannel
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetCurrentCellId() uint32 {
	if m != nil {
		return m.CurrentCellId
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetSecondsUntilSlotEnd() float32 {
	if m != nil {
		return m.SecondsUntilSlotEnd
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetWbRssiPeakMagDb() float32 {
	if m != nil {
		return m.WbRssiPeakMagDb
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetPopPingDropRate() float32 {
	if m != nil {
		return m.PopPingDropRate
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetSnrDb() float32 {
	if m != nil {
		return m.SnrDb
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetL1SnrAvgDb() float32 {
	if m != nil {
		return m.L1SnrAvgDb
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetL1SnrMinDb() float32 {
	if m != nil {
		return m.L1SnrMinDb
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetL1SnrMaxDb() float32 {
	if m != nil {
		return m.L1SnrMaxDb
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetLmacSatelliteId() uint32 {
	if m != nil {
		return m.LmacSatelliteId
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetTargetSatelliteId() uint32 {
	if m != nil {
		return m.TargetSatelliteId
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetGrantMcs() uint32 {
	if m != nil {
		return m.GrantMcs
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetGrantSymbolsAvg() float32 {
	if m != nil {
		return m.GrantSymbolsAvg
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetDedGrant() uint32 {
	if m != nil {
		return m.DedGrant
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetMobilityProactiveSlotChange() uint32 {
	if m != nil {
		return m.MobilityProactiveSlotChange
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetMobilityReactiveSlotChange() uint32 {
	if m != nil {
		return m.MobilityReactiveSlotChange
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetRfpTotalSynFailed() uint32 {
	if m != nil {
		return m.RfpTotalSynFailed
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetNumOutOfSeq() uint32 {
	if m != nil {
		return m.NumOutOfSeq
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetNumUlmapDrop() uint32 {
	if m != nil {
		return m.NumUlmapDrop
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetCurrentSecondsOfSchedule() float32 {
	if m != nil {
		return m.CurrentSecondsOfSchedule
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetSendLabelSwitchToGroundFailedCalls() uint32 {
	if m != nil {
		return m.SendLabelSwitchToGroundFailedCalls
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetEmaVelocityX() float64 {
	if m != nil {
		return m.EmaVelocityX
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetEmaVelocityY() float64 {
	if m != nil {
		return m.EmaVelocityY
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetEmaVelocityZ() float64 {
	if m != nil {
		return m.EmaVelocityZ
	}
	return 0
}

func (m *TransceiverGetTelemetryResponse) GetCeRssiDb() float32 {
	if m != nil {
		return m.CeRssiDb
	}
	return 0
}

func init() {
	proto.RegisterEnum("SpaceX.API.Device.TransceiverModulatorState", TransceiverModulatorState_name, TransceiverModulatorState_value)
	proto.RegisterEnum("SpaceX.API.Device.TransceiverTxRxState", TransceiverTxRxState_name, TransceiverTxRxState_value)
	proto.RegisterEnum("SpaceX.API.Device.TransceiverTransmitBlankingState", TransceiverTransmitBlankingState_name, TransceiverTransmitBlankingState_value)
	proto.RegisterType((*TransceiverIFLoopbackTestRequest)(nil), "SpaceX.API.Device.TransceiverIFLoopbackTestRequest")
	proto.RegisterType((*TransceiverIFLoopbackTestResponse)(nil), "SpaceX.API.Device.TransceiverIFLoopbackTestResponse")
	proto.RegisterType((*TransceiverGetStatusRequest)(nil), "SpaceX.API.Device.TransceiverGetStatusRequest")
	proto.RegisterType((*TransceiverGetStatusResponse)(nil), "SpaceX.API.Device.TransceiverGetStatusResponse")
	proto.RegisterType((*TransceiverFaults)(nil), "SpaceX.API.Device.TransceiverFaults")
	proto.RegisterType((*TransceiverGetTelemetryRequest)(nil), "SpaceX.API.Device.TransceiverGetTelemetryRequest")
	proto.RegisterType((*TransceiverGetTelemetryResponse)(nil), "SpaceX.API.Device.TransceiverGetTelemetryResponse")
}

func init() {
	proto.RegisterFile("spacex/api/device/transceiver.proto", fileDescriptor_528fde07cc1e136a)
}

var fileDescriptor_528fde07cc1e136a = []byte{
	// 1357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdb, 0x52, 0xdb, 0x48,
	0x1a, 0xc7, 0x17, 0x72, 0xc0, 0x69, 0x30, 0xd8, 0x0d, 0x24, 0x0a, 0x21, 0x6c, 0xd6, 0x9b, 0xad,
	0x64, 0xb3, 0x59, 0xd8, 0x84, 0xbd, 0xda, 0xaa, 0xdd, 0x2a, 0xc0, 0x24, 0xeb, 0x1a, 0x0c, 0x94,
	0x6c, 0x32, 0x4c, 0x2e, 0xa6, 0xab, 0x25, 0x7d, 0x36, 0x2a, 0x5a, 0xdd, 0x4a, 0x77, 0xcb, 0x31,
	0x73, 0x35, 0xe7, 0x87, 0x99, 0xe7, 0x98, 0x97, 0x98, 0x47, 0x98, 0xcc, 0xf9, 0x7c, 0x3e, 0x54,
	0x77, 0x4b, 0x60, 0xc7, 0x24, 0x53, 0x33, 0x77, 0xf6, 0xf7, 0xff, 0xfd, 0xff, 0xad, 0xaf, 0xf5,
	0xb5, 0x24, 0xf4, 0x57, 0x95, 0xd2, 0x10, 0xfa, 0x2b, 0x34, 0x8d, 0x57, 0x22, 0xe8, 0xc5, 0x21,
	0xac, 0x68, 0x49, 0xb9, 0x0a, 0x21, 0xee, 0x81, 0x5c, 0x4e, 0xa5, 0xd0, 0x02, 0x57, 0x5b, 0x06,
	0xda, 0x5f, 0x5e, 0xdb, 0x6d, 0x2c, 0xd7, 0x2d, 0xb4, 0xb0, 0x34, 0xea, 0x0b, 0x45, 0x92, 0x08,
	0xee, 0x2c, 0x0b, 0x8b, 0xa3, 0x7a, 0x14, 0xab, 0x03, 0xa7, 0xd6, 0x76, 0xd1, 0xb5, 0xf6, 0xc9,
	0x2a, 0x8d, 0x7b, 0x5b, 0x42, 0xa4, 0x01, 0x0d, 0x0f, 0xdb, 0xa0, 0xb4, 0x0f, 0x8f, 0x32, 0x50,
	0x1a, 0xdf, 0x46, 0x18, 0x38, 0x0d, 0x18, 0x90, 0xb8, 0x43, 0x58, 0x0e, 0x78, 0x63, 0xd7, 0xc6,
	0x6e, 0x96, 0xfc, 0x8a, 0x53, 0x1a, 0x9d, 0xc2, 0x58, 0x7b, 0x77, 0x0c, 0xfd, 0xe5, 0x39, 0x91,
	0x2a, 0x15, 0x5c, 0x01, 0xbe, 0x85, 0xaa, 0x01, 0xc8, 0xe3, 0x34, 0xa2, 0x41, 0x69, 0x1b, 0x39,
	0xee, 0xcf, 0x04, 0x20, 0x07, 0x3d, 0x86, 0x55, 0xfc, 0x69, 0x76, 0xdc, 0xb1, 0x8a, 0x0f, 0xb3,
	0xb7, 0x11, 0x96, 0x4a, 0xc5, 0x4f, 0xc1, 0x67, 0x2c, 0x5c, 0x31, 0xca, 0x10, 0x7d, 0x19, 0x95,
	0x52, 0xc6, 0x08, 0x13, 0xe1, 0xa1, 0x77, 0xd6, 0xf6, 0x33, 0x91, 0x32, 0xb6, 0x25, 0xc2, 0xc3,
	0xda, 0x55, 0x74, 0x65, 0xa0, 0x8b, 0xfb, 0xa0, 0x5b, 0x9a, 0xea, 0x4c, 0xe5, 0x7b, 0x52, 0x7b,
	0x72, 0x16, 0x2d, 0x9e, 0xae, 0xe7, 0x0d, 0x36, 0xd0, 0x85, 0x44, 0x44, 0x44, 0x69, 0xaa, 0xc1,
	0x36, 0x36, 0x7d, 0xf7, 0xf6, 0xf2, 0xc8, 0xdd, 0x5b, 0x1e, 0xc8, 0x68, 0x8a, 0x28, 0x63, 0x54,
	0x0b, 0x69, 0x92, 0xc0, 0x2f, 0x25, 0x22, 0xb2, 0xbf, 0x70, 0x13, 0x4d, 0x46, 0x70, 0x12, 0x36,
	0xfe, 0x07, 0xc2, 0x90, 0x0d, 0x70, 0x71, 0xeb, 0xa8, 0xa4, 0xfb, 0x79, 0xd6, 0x19, 0x9b, 0x75,
	0xe3, 0xf9, 0x59, 0xed, 0xbe, 0xdf, 0x77, 0x31, 0x13, 0xba, 0x7f, 0x9c, 0x21, 0x8b, 0x8c, 0xb3,
	0xbf, 0x33, 0x43, 0xe6, 0x19, 0xab, 0xe8, 0x9c, 0x0b, 0xf8, 0x78, 0xc2, 0x26, 0x2c, 0x9e, 0x92,
	0x50, 0x8f, 0xd5, 0x81, 0xb3, 0x39, 0x16, 0xff, 0x17, 0x9d, 0xef, 0xd0, 0x8c, 0x69, 0xe5, 0x7d,
	0x62, 0x5c, 0x93, 0x77, 0xaf, 0x3f, 0x7f, 0xdd, 0x7b, 0x16, 0xf6, 0x73, 0x13, 0x66, 0xe8, 0x92,
	0x3d, 0x54, 0x49, 0xac, 0x49, 0xc0, 0x28, 0x3f, 0x8c, 0x79, 0x37, 0x6f, 0xe3, 0x53, 0x77, 0x15,
	0xab, 0xbf, 0xd1, 0x47, 0xee, 0x5e, 0xcf, 0xcd, 0xee, 0xe2, 0xe6, 0xf5, 0x69, 0x65, 0x7c, 0x03,
	0xcd, 0x24, 0x22, 0x82, 0x84, 0x50, 0x15, 0x87, 0x44, 0x43, 0x92, 0x7a, 0x9f, 0x4d, 0xd8, 0x51,
	0x2c, 0xdb, 0xfa, 0x9a, 0x8a, 0xc3, 0x36, 0x24, 0x29, 0xbe, 0x8a, 0x90, 0xee, 0x9b, 0xd3, 0x65,
	0x99, 0xcf, 0x1d, 0x53, 0xd2, 0xfd, 0x46, 0xc7, 0xc8, 0xb5, 0x77, 0xc6, 0x50, 0x75, 0xa4, 0x27,
	0xfc, 0x1f, 0xb4, 0x20, 0x7a, 0x20, 0xad, 0x87, 0x0c, 0xac, 0x63, 0x5b, 0xcd, 0x8f, 0xe7, 0x45,
	0x43, 0x98, 0x8c, 0x66, 0xb1, 0x9e, 0x35, 0xe3, 0x15, 0x34, 0x77, 0xe2, 0x4d, 0xc3, 0x80, 0xe6,
	0xae, 0x71, 0xeb, 0xaa, 0x16, 0xae, 0xdd, 0x30, 0xa0, 0xce, 0x70, 0x13, 0x55, 0xa2, 0x90, 0xf4,
	0x04, 0xd3, 0xb4, 0x0b, 0x39, 0x7c, 0xc6, 0xc2, 0xd3, 0x51, 0xf8, 0xc0, 0x95, 0x2d, 0x59, 0xbb,
	0x86, 0x96, 0x86, 0x0f, 0x46, 0x1b, 0x18, 0x24, 0xa0, 0xe5, 0x51, 0x71, 0x76, 0xde, 0x43, 0xe8,
	0xcf, 0xcf, 0x44, 0xf2, 0xe3, 0xb3, 0x8a, 0xe6, 0x29, 0xd7, 0xc0, 0x39, 0x25, 0xa9, 0x88, 0xb9,
	0x36, 0xf7, 0xc9, 0xf4, 0xe8, 0xbd, 0x6f, 0x36, 0xa7, 0xec, 0xcf, 0xe6, 0xea, 0x6e, 0x2e, 0x9a,
	0xfe, 0xf0, 0x75, 0x54, 0x3e, 0x36, 0xc5, 0x3a, 0x3c, 0xf0, 0x9e, 0xb8, 0x9d, 0x9c, 0x2a, 0x60,
	0x53, 0xc4, 0x35, 0x54, 0xfc, 0x27, 0x52, 0x30, 0xe6, 0x7d, 0xe0, 0xa0, 0xc9, 0xbc, 0xe8, 0x0b,
	0xc6, 0xf0, 0xdf, 0x51, 0xe5, 0x98, 0xe9, 0x13, 0x7d, 0x00, 0x9a, 0x7a, 0x1f, 0x3a, 0x6e, 0xba,
	0xe0, 0xfa, 0x6d, 0x53, 0xc6, 0x77, 0xd0, 0x5c, 0x81, 0x6a, 0x99, 0x01, 0x39, 0x00, 0x1a, 0xc5,
	0xbc, 0xeb, 0x7d, 0xe4, 0x70, 0x9c, 0x8b, 0x6d, 0x99, 0xc1, 0xff, 0x9d, 0x84, 0x97, 0x10, 0x92,
	0x7d, 0x12, 0x1e, 0x50, 0xce, 0x81, 0xb9, 0xf1, 0x2f, 0xfb, 0x17, 0x64, 0x7f, 0xc3, 0x55, 0xcc,
	0xdc, 0x84, 0x99, 0x94, 0xc0, 0x35, 0x09, 0x81, 0x31, 0x12, 0x47, 0x6e, 0xda, 0xcb, 0x7e, 0x39,
	0xaf, 0x6f, 0x00, 0x63, 0x8d, 0x08, 0xff, 0x1b, 0x5d, 0x54, 0x10, 0x0a, 0x1e, 0x29, 0x92, 0x71,
	0x1d, 0x33, 0xa2, 0x98, 0xd0, 0x04, 0x78, 0xe4, 0xa6, 0x79, 0xdc, 0x9f, 0xcd, 0xe5, 0x3d, 0xa3,
	0xb6, 0x98, 0xd0, 0x9b, 0x3c, 0xc2, 0xff, 0x44, 0xb3, 0x8f, 0x03, 0x62, 0x1f, 0x93, 0x29, 0xd0,
	0x43, 0x92, 0xd0, 0x2e, 0x89, 0x82, 0x62, 0x34, 0x67, 0x1e, 0x07, 0xbe, 0x52, 0xf1, 0x2e, 0xd0,
	0xc3, 0x26, 0xed, 0xd6, 0x03, 0xf3, 0x48, 0x4d, 0x45, 0x4a, 0x52, 0x73, 0x0b, 0x22, 0x29, 0x52,
	0x22, 0xcd, 0x71, 0xc9, 0x87, 0x74, 0x26, 0x15, 0xe9, 0x6e, 0xcc, 0xbb, 0x75, 0x29, 0x52, 0xdf,
	0xcc, 0xfc, 0x45, 0x74, 0xde, 0x3c, 0xac, 0xa3, 0xc0, 0xfb, 0xc2, 0x11, 0xe7, 0x14, 0x97, 0xf5,
	0x00, 0xd7, 0x50, 0x99, 0xdd, 0x21, 0x46, 0xa2, 0x3d, 0xbb, 0xdc, 0x97, 0x4e, 0x46, 0xec, 0x4e,
	0x8b, 0xcb, 0xb5, 0x5e, 0x77, 0x88, 0x49, 0x62, 0x6e, 0x98, 0xaf, 0x06, 0x99, 0x66, 0xcc, 0x87,
	0x19, 0xda, 0x37, 0xcc, 0xd7, 0x43, 0x0c, 0xed, 0xd7, 0x03, 0xfc, 0x0f, 0x54, 0x65, 0x09, 0x0d,
	0x89, 0xa2, 0x1a, 0x18, 0x8b, 0x35, 0x98, 0x1d, 0xfc, 0xc6, 0xed, 0xe0, 0x8c, 0x51, 0x5a, 0x85,
	0xd0, 0x88, 0xf0, 0x0a, 0x9a, 0xd5, 0x54, 0x76, 0x41, 0x0f, 0xe3, 0xdf, 0x3a, 0xbc, 0xea, 0xb4,
	0x41, 0xc3, 0x22, 0xba, 0xd0, 0x95, 0x94, 0x6b, 0x92, 0x84, 0xca, 0xfb, 0xce, 0x61, 0x25, 0x5b,
	0x69, 0x86, 0xca, 0xac, 0xed, 0x54, 0x75, 0x94, 0x04, 0x82, 0x29, 0xd3, 0xae, 0xf7, 0x7d, 0xbe,
	0x59, 0x56, 0x69, 0x39, 0x61, 0xad, 0xd7, 0x35, 0x51, 0x11, 0x44, 0xc4, 0x96, 0xbd, 0x1f, 0xf2,
	0xa8, 0x08, 0xa2, 0xfb, 0xa6, 0x80, 0xeb, 0x68, 0x29, 0x11, 0x41, 0xcc, 0x62, 0x7d, 0x44, 0x52,
	0x29, 0x68, 0xa8, 0xe3, 0x1e, 0xb8, 0x5b, 0x6c, 0x66, 0xa7, 0x0b, 0xde, 0x8f, 0xce, 0x72, 0xa5,
	0xc0, 0x76, 0x0b, 0xca, 0xdc, 0xea, 0x0d, 0xcb, 0xe0, 0x75, 0x74, 0xf5, 0x38, 0x45, 0xc2, 0x29,
	0x21, 0x3f, 0xb9, 0x90, 0x85, 0x82, 0xf2, 0x61, 0x24, 0xe3, 0x5f, 0x68, 0x4e, 0x76, 0x52, 0xa2,
	0x85, 0xa6, 0x8c, 0xa8, 0x23, 0x4e, 0x3a, 0x34, 0x66, 0x10, 0x79, 0x3f, 0xe7, 0x9b, 0x24, 0x3b,
	0x69, 0xdb, 0x68, 0xad, 0x23, 0x7e, 0xcf, 0x2a, 0xf8, 0x3a, 0x9a, 0xe6, 0x59, 0x42, 0x44, 0xa6,
	0x89, 0xe8, 0x10, 0x05, 0x8f, 0xbc, 0x5f, 0x1c, 0x3b, 0xc9, 0xb3, 0x64, 0x27, 0xd3, 0x3b, 0x9d,
	0x16, 0x3c, 0xc2, 0x7f, 0x73, 0x54, 0xc6, 0x12, 0x9a, 0xda, 0xd9, 0xf2, 0x5e, 0x2d, 0x59, 0x6a,
	0x8a, 0x67, 0xc9, 0x9e, 0xa9, 0x9a, 0xb9, 0xc2, 0xff, 0x43, 0x57, 0x8a, 0xf3, 0x50, 0x8c, 0xbb,
	0x09, 0x0d, 0x0f, 0x20, 0xca, 0x18, 0x78, 0xaf, 0x95, 0xec, 0xee, 0x7a, 0x39, 0xd3, 0x72, 0xc8,
	0x4e, 0xa7, 0x95, 0x03, 0x78, 0x0f, 0xdd, 0x54, 0xc0, 0x23, 0xc2, 0x68, 0x00, 0x8c, 0xa8, 0xc7,
	0xe6, 0x31, 0x40, 0xb4, 0x20, 0x5d, 0x29, 0x32, 0x1e, 0xe5, 0xcd, 0x90, 0x90, 0x32, 0xa6, 0xbc,
	0xd7, 0xdd, 0x05, 0xd4, 0x8c, 0x61, 0xcb, 0xf0, 0x2d, 0x8b, 0xb7, 0xc5, 0x7d, 0x0b, 0xbb, 0xf6,
	0x36, 0x0c, 0x6a, 0xae, 0x1e, 0x12, 0x4a, 0x7a, 0xc0, 0x44, 0x68, 0x76, 0xb7, 0xef, 0xbd, 0x61,
	0xcc, 0x63, 0xfe, 0x14, 0x24, 0xf4, 0x41, 0x5e, 0xdd, 0x1f, 0xc1, 0x8e, 0xbc, 0x37, 0x47, 0xb1,
	0x97, 0x46, 0xb0, 0x57, 0xbc, 0xb7, 0x46, 0xb1, 0x87, 0xe6, 0x55, 0x11, 0x82, 0x3b, 0xbc, 0x51,
	0xe0, 0xbd, 0xed, 0x5a, 0x2f, 0x85, 0x60, 0xce, 0x6c, 0x3d, 0xb8, 0xf5, 0x32, 0xba, 0xfc, 0xcc,
	0xaf, 0x00, 0x3c, 0x87, 0x2a, 0xcd, 0x9d, 0x7a, 0xab, 0xbd, 0xd6, 0xde, 0x24, 0x7b, 0xdb, 0x2f,
	0x6c, 0xef, 0xbc, 0xb8, 0x5d, 0xf9, 0xd3, 0x50, 0x75, 0x73, 0x7b, 0x6d, 0x7d, 0x6b, 0xb3, 0x5e,
	0x19, 0xc3, 0xf3, 0xa8, 0x7a, 0x5c, 0xad, 0x37, 0x5a, 0xae, 0x3c, 0x7e, 0xab, 0x89, 0xe6, 0x4e,
	0x7b, 0xab, 0xe3, 0x0a, 0x9a, 0x6a, 0xef, 0xfb, 0xfb, 0x03, 0xb1, 0x45, 0xe5, 0x24, 0xb2, 0x8a,
	0xca, 0xb6, 0x32, 0x10, 0xd7, 0x1a, 0xfa, 0xfc, 0x3c, 0xf5, 0xe5, 0x8a, 0xa7, 0x11, 0x6a, 0xaf,
	0x0f, 0x04, 0xbb, 0xff, 0x27, 0xb1, 0x33, 0x68, 0xb2, 0xbd, 0x3e, 0x10, 0xba, 0x7e, 0xe9, 0xe1,
	0xbc, 0xfb, 0xe6, 0x5d, 0x0e, 0x45, 0x32, 0xf0, 0xdd, 0x1b, 0x9c, 0xb7, 0xdf, 0xbc, 0xab, 0xbf,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x08, 0x0f, 0xee, 0xcc, 0x6b, 0x0b, 0x00, 0x00,
}

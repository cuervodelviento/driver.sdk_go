package config

type NetsocsConfigKey string

const DELETE_USER NetsocsConfigKey = "deleteUser"
const GET_CHANNELS NetsocsConfigKey = "getChannels"
const GET_USERS NetsocsConfigKey = "getUsers"
const SET_USERS NetsocsConfigKey = "setUsers"
const ACTION_LISTEN_EVENTS NetsocsConfigKey = "actionListenEvent"

const ACTION_OUTPUT_PULSE NetsocsConfigKey = "actionOutputPulse"
const ACTION_PING_DEVICE NetsocsConfigKey = "actionPingDevice"
const ACTION_RESTART_DEVICE NetsocsConfigKey = "actionRestart"
const ACTION_STOP_LISTEN_EVENT NetsocsConfigKey = "actionStopListenEvent"
const SET_DELETE_STORAGE NetsocsConfigKey = "setDeleteStorage"
const SET_DELETE_USER NetsocsConfigKey = "setDeleteUser"
const SET_CARD_TO_PERSON_AC NetsocsConfigKey = "setCardToPersonAC"
const SET_FACE_TO_PERSON_AC NetsocsConfigKey = "setFaceToPersonAC"
const SET_ADD_PERSON_TO_AC NetsocsConfigKey = "setAddPersonToAC"
const SET_DEL_PERSON_TO_AC NetsocsConfigKey = "setDelPersonToAC"
const GET_RECORDING_SOURCE NetsocsConfigKey = "getRecordingSource"
const ACTION_ZOOM NetsocsConfigKey = "actionZoom"
const GET_AVAILABLE_OUTPUTS NetsocsConfigKey = "getAvailableOutputs"
const GET_AVAILABLE_SPEAKERS NetsocsConfigKey = "getAvailableSpeakers"
const GET_STORAGES NetsocsConfigKey = "getStorages"
const GET_CURRENT_VIDEO_RESOLUTION_BY_CHANNEL = "getCurrentVideoResolutionByChannel"
const GET_AVAILABLE_VIDEO_RESOLUTIONS NetsocsConfigKey = "getAvailableVideoResolutions"
const SET_VIDEO_RESOLUTION NetsocsConfigKey = "setVideoResolution"
const SET_VIDEO_IN_BLACK_AND_WHITE NetsocsConfigKey = "setVideoInBlackAndWhite"
const GET_VIDEO_IN_BLACK_AND_WHITE_STATUS NetsocsConfigKey = "getVideoInBlackAndWhiteStatus"
const SET_MIRROR_VIDEO NetsocsConfigKey = "setMirrorVideo"
const GET_MIRROR_VIDEO_STATUS NetsocsConfigKey = "getMirrorVideoStatus"
const SET_FLIP_VIDEO NetsocsConfigKey = "setFlipVideo"
const GET_FLIP_VIDEO_STATUS NetsocsConfigKey = "getFlipVideoStatus"
const GET_HEATMAP_IMAGE NetsocsConfigKey = "getHeatmapImage"
const SET_BLOCK_PERSON_TO_AC NetsocsConfigKey = "setBlockPersonToAC"
const DELETE_ALL_PEOPLE_AC NetsocsConfigKey = "deleteAllPeopleAC"
const GET_INPUTS NetsocsConfigKey = "getInputs"
const GET_MICROPHONES NetsocsConfigKey = "getMicrophones"
const GET_FTP_INFO NetsocsConfigKey = "getFtpInfo"
const SET_FTP_INFO NetsocsConfigKey = "setFtpInfo"
const ACTION_PLAY_AUDIO_CLIP NetsocsConfigKey = "actionPlayAudioClip"
const GET_ALARM_PARTITIONS NetsocsConfigKey = "getAlarmPartitions"
const SET_ALARM_PARTITION NetsocsConfigKey = "setAlarmPartition"
const GET_ALARM_ZONES NetsocsConfigKey = "getAlarmZones"
const SET_ALARM_ZONE NetsocsConfigKey = "setAlarmZone"
const GET_ALARM_ZONE_STATUS NetsocsConfigKey = "getAlarmZoneStatus"
const SET_ADD_ALARM_USER NetsocsConfigKey = "setAddAlarmUser"
const SET_ALARM_USER NetsocsConfigKey = "setAlarmUser"
const GET_ALARM_USERS NetsocsConfigKey = "getAlarmUsers"

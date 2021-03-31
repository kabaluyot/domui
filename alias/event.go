package alias

import "github.com/reusee/domui"

// https://developer.mozilla.org/en-US/docs/Web/Events

var (
	EDOMActivate                  = domui.MakeEventFunc("DOMActivate")
	EDOMContentLoaded             = domui.MakeEventFunc("DOMContentLoaded")
	EDOMMouseScroll               = domui.MakeEventFunc("DOMMouseScroll")
	EMSGestureChange              = domui.MakeEventFunc("MSGestureChange")
	EMSGestureEnd                 = domui.MakeEventFunc("MSGestureEnd")
	EMSGestureHold                = domui.MakeEventFunc("MSGestureHold")
	EMSGestureStart               = domui.MakeEventFunc("MSGestureStart")
	EMSGestureTap                 = domui.MakeEventFunc("MSGestureTap")
	EMSInertiaStart               = domui.MakeEventFunc("MSInertiaStart")
	EMSManipulationStateChanged   = domui.MakeEventFunc("MSManipulationStateChanged")
	Eabort                        = domui.MakeEventFunc("abort")
	Eactivate                     = domui.MakeEventFunc("activate")
	Eaddstream                    = domui.MakeEventFunc("addstream")
	Eaddtrack                     = domui.MakeEventFunc("addtrack")
	Eafterprint                   = domui.MakeEventFunc("afterprint")
	Eafterscriptexecute           = domui.MakeEventFunc("afterscriptexecute")
	Eanimationcancel              = domui.MakeEventFunc("animationcancel")
	Eanimationend                 = domui.MakeEventFunc("animationend")
	Eanimationiteration           = domui.MakeEventFunc("animationiteration")
	Eanimationstart               = domui.MakeEventFunc("animationstart")
	Eappinstalled                 = domui.MakeEventFunc("appinstalled")
	Eaudioend                     = domui.MakeEventFunc("audioend")
	Eaudioprocess                 = domui.MakeEventFunc("audioprocess")
	Eaudiostart                   = domui.MakeEventFunc("audiostart")
	Eauxclick                     = domui.MakeEventFunc("auxclick")
	Ebeforeinput                  = domui.MakeEventFunc("beforeinput")
	Ebeforeprint                  = domui.MakeEventFunc("beforeprint")
	Ebeforescriptexecute          = domui.MakeEventFunc("beforescriptexecute")
	Ebeforeunload                 = domui.MakeEventFunc("beforeunload")
	EbeginEvent                   = domui.MakeEventFunc("beginEvent")
	Eblocked                      = domui.MakeEventFunc("blocked")
	Eblur                         = domui.MakeEventFunc("blur")
	Eboundary                     = domui.MakeEventFunc("boundary")
	Ebufferedamountlow            = domui.MakeEventFunc("bufferedamountlow")
	Ecancel                       = domui.MakeEventFunc("cancel")
	Ecanplay                      = domui.MakeEventFunc("canplay")
	Ecanplaythrough               = domui.MakeEventFunc("canplaythrough")
	Echange                       = domui.MakeEventFunc("change")
	Eclick                        = domui.MakeEventFunc("click")
	Eclose                        = domui.MakeEventFunc("close")
	Eclosing                      = domui.MakeEventFunc("closing")
	Ecomplete                     = domui.MakeEventFunc("complete")
	Ecompositionend               = domui.MakeEventFunc("compositionend")
	Ecompositionstart             = domui.MakeEventFunc("compositionstart")
	Ecompositionupdate            = domui.MakeEventFunc("compositionupdate")
	Econnect                      = domui.MakeEventFunc("connect")
	Econnectionstatechange        = domui.MakeEventFunc("connectionstatechange")
	Econtentdelete                = domui.MakeEventFunc("contentdelete")
	Econtextmenu                  = domui.MakeEventFunc("contextmenu")
	Ecopy                         = domui.MakeEventFunc("copy")
	Ecuechange                    = domui.MakeEventFunc("cuechange")
	Ecut                          = domui.MakeEventFunc("cut")
	Edatachannel                  = domui.MakeEventFunc("datachannel")
	Edblclick                     = domui.MakeEventFunc("dblclick")
	Edevicechange                 = domui.MakeEventFunc("devicechange")
	Edevicemotion                 = domui.MakeEventFunc("devicemotion")
	Edeviceorientation            = domui.MakeEventFunc("deviceorientation")
	Edrag                         = domui.MakeEventFunc("drag")
	Edragend                      = domui.MakeEventFunc("dragend")
	Edragenter                    = domui.MakeEventFunc("dragenter")
	Edragleave                    = domui.MakeEventFunc("dragleave")
	Edragover                     = domui.MakeEventFunc("dragover")
	Edragstart                    = domui.MakeEventFunc("dragstart")
	Edrop                         = domui.MakeEventFunc("drop")
	Edurationchange               = domui.MakeEventFunc("durationchange")
	Eemptied                      = domui.MakeEventFunc("emptied")
	Eend                          = domui.MakeEventFunc("end")
	EendEvent                     = domui.MakeEventFunc("endEvent")
	Eended                        = domui.MakeEventFunc("ended")
	Eenterpictureinpicture        = domui.MakeEventFunc("enterpictureinpicture")
	Eerror                        = domui.MakeEventFunc("error")
	Efocus                        = domui.MakeEventFunc("focus")
	Efocusin                      = domui.MakeEventFunc("focusin")
	Efocusout                     = domui.MakeEventFunc("focusout")
	Eformdata                     = domui.MakeEventFunc("formdata")
	Efullscreenchange             = domui.MakeEventFunc("fullscreenchange")
	Efullscreenerror              = domui.MakeEventFunc("fullscreenerror")
	Egamepadconnected             = domui.MakeEventFunc("gamepadconnected")
	Egamepaddisconnected          = domui.MakeEventFunc("gamepaddisconnected")
	Egatheringstatechange         = domui.MakeEventFunc("gatheringstatechange")
	Egesturechange                = domui.MakeEventFunc("gesturechange")
	Egestureend                   = domui.MakeEventFunc("gestureend")
	Egesturestart                 = domui.MakeEventFunc("gesturestart")
	Egotpointercapture            = domui.MakeEventFunc("gotpointercapture")
	Ehashchange                   = domui.MakeEventFunc("hashchange")
	Eicecandidate                 = domui.MakeEventFunc("icecandidate")
	Eicecandidateerror            = domui.MakeEventFunc("icecandidateerror")
	Eiceconnectionstatechange     = domui.MakeEventFunc("iceconnectionstatechange")
	Eicegatheringstatechange      = domui.MakeEventFunc("icegatheringstatechange")
	Eidentityresult               = domui.MakeEventFunc("identityresult")
	Eidpassertionerror            = domui.MakeEventFunc("idpassertionerror")
	Eidpvalidationerror           = domui.MakeEventFunc("idpvalidationerror")
	Einput                        = domui.MakeEventFunc("input")
	Einputsourceschange           = domui.MakeEventFunc("inputsourceschange")
	Einstall                      = domui.MakeEventFunc("install")
	Einvalid                      = domui.MakeEventFunc("invalid")
	Ekeydown                      = domui.MakeEventFunc("keydown")
	Ekeypress                     = domui.MakeEventFunc("keypress")
	Ekeyup                        = domui.MakeEventFunc("keyup")
	Elanguagechange               = domui.MakeEventFunc("languagechange")
	Eleavepictureinpicture        = domui.MakeEventFunc("leavepictureinpicture")
	Eload                         = domui.MakeEventFunc("load")
	Eloadeddata                   = domui.MakeEventFunc("loadeddata")
	Eloadedmetadata               = domui.MakeEventFunc("loadedmetadata")
	Eloadend                      = domui.MakeEventFunc("loadend")
	Eloadstart                    = domui.MakeEventFunc("loadstart")
	Elostpointercapture           = domui.MakeEventFunc("lostpointercapture")
	Emark                         = domui.MakeEventFunc("mark")
	Emerchantvalidation           = domui.MakeEventFunc("merchantvalidation")
	Emessage                      = domui.MakeEventFunc("message")
	Emessageerror                 = domui.MakeEventFunc("messageerror")
	Emousedown                    = domui.MakeEventFunc("mousedown")
	Emouseenter                   = domui.MakeEventFunc("mouseenter")
	Emouseleave                   = domui.MakeEventFunc("mouseleave")
	Emousemove                    = domui.MakeEventFunc("mousemove")
	Emouseout                     = domui.MakeEventFunc("mouseout")
	Emouseover                    = domui.MakeEventFunc("mouseover")
	Emouseup                      = domui.MakeEventFunc("mouseup")
	Emousewheel                   = domui.MakeEventFunc("mousewheel")
	EmsContentZoom                = domui.MakeEventFunc("msContentZoom")
	Emute                         = domui.MakeEventFunc("mute")
	Enegotiationneeded            = domui.MakeEventFunc("negotiationneeded")
	Enomatch                      = domui.MakeEventFunc("nomatch")
	Enotificationclick            = domui.MakeEventFunc("notificationclick")
	Eoffline                      = domui.MakeEventFunc("offline")
	Eonline                       = domui.MakeEventFunc("online")
	Eopen                         = domui.MakeEventFunc("open")
	Eorientationchange            = domui.MakeEventFunc("orientationchange")
	Eoverflow                     = domui.MakeEventFunc("overflow")
	Epagehide                     = domui.MakeEventFunc("pagehide")
	Epageshow                     = domui.MakeEventFunc("pageshow")
	Epaste                        = domui.MakeEventFunc("paste")
	Epause                        = domui.MakeEventFunc("pause")
	Epayerdetailchange            = domui.MakeEventFunc("payerdetailchange")
	Epaymentmethodchange          = domui.MakeEventFunc("paymentmethodchange")
	Epeeridentity                 = domui.MakeEventFunc("peeridentity")
	Eplay                         = domui.MakeEventFunc("play")
	Eplaying                      = domui.MakeEventFunc("playing")
	Epointercancel                = domui.MakeEventFunc("pointercancel")
	Epointerdown                  = domui.MakeEventFunc("pointerdown")
	Epointerenter                 = domui.MakeEventFunc("pointerenter")
	Epointerleave                 = domui.MakeEventFunc("pointerleave")
	Epointerlockchange            = domui.MakeEventFunc("pointerlockchange")
	Epointerlockerror             = domui.MakeEventFunc("pointerlockerror")
	Epointermove                  = domui.MakeEventFunc("pointermove")
	Epointerout                   = domui.MakeEventFunc("pointerout")
	Epointerover                  = domui.MakeEventFunc("pointerover")
	Epointerup                    = domui.MakeEventFunc("pointerup")
	Epopstate                     = domui.MakeEventFunc("popstate")
	Eprogress                     = domui.MakeEventFunc("progress")
	Epush                         = domui.MakeEventFunc("push")
	Epushsubscriptionchange       = domui.MakeEventFunc("pushsubscriptionchange")
	Eratechange                   = domui.MakeEventFunc("ratechange")
	Ereadystatechange             = domui.MakeEventFunc("readystatechange")
	Erejectionhandled             = domui.MakeEventFunc("rejectionhandled")
	EremoveTrack                  = domui.MakeEventFunc("removeTrack")
	Eremovestream                 = domui.MakeEventFunc("removestream")
	Eremovetrack                  = domui.MakeEventFunc("removetrack")
	ErepeatEvent                  = domui.MakeEventFunc("repeatEvent")
	Ereset                        = domui.MakeEventFunc("reset")
	Eresize                       = domui.MakeEventFunc("resize")
	Eresourcetimingbufferfull     = domui.MakeEventFunc("resourcetimingbufferfull")
	Eresult                       = domui.MakeEventFunc("result")
	Eresume                       = domui.MakeEventFunc("resume")
	Escroll                       = domui.MakeEventFunc("scroll")
	Esearch                       = domui.MakeEventFunc("search")
	Eseeked                       = domui.MakeEventFunc("seeked")
	Eseeking                      = domui.MakeEventFunc("seeking")
	Eselect                       = domui.MakeEventFunc("select")
	Eselectedcandidatepairchange  = domui.MakeEventFunc("selectedcandidatepairchange")
	Eselectend                    = domui.MakeEventFunc("selectend")
	Eselectionchange              = domui.MakeEventFunc("selectionchange")
	Eselectstart                  = domui.MakeEventFunc("selectstart")
	Eshippingaddresschange        = domui.MakeEventFunc("shippingaddresschange")
	Eshippingoptionchange         = domui.MakeEventFunc("shippingoptionchange")
	Eshow                         = domui.MakeEventFunc("show")
	Esignalingstatechange         = domui.MakeEventFunc("signalingstatechange")
	Eslotchange                   = domui.MakeEventFunc("slotchange")
	Esoundend                     = domui.MakeEventFunc("soundend")
	Esoundstart                   = domui.MakeEventFunc("soundstart")
	Espeechend                    = domui.MakeEventFunc("speechend")
	Espeechstart                  = domui.MakeEventFunc("speechstart")
	Esqueeze                      = domui.MakeEventFunc("squeeze")
	Esqueezeend                   = domui.MakeEventFunc("squeezeend")
	Esqueezestart                 = domui.MakeEventFunc("squeezestart")
	Estalled                      = domui.MakeEventFunc("stalled")
	Estart                        = domui.MakeEventFunc("start")
	Estatechange                  = domui.MakeEventFunc("statechange")
	Estorage                      = domui.MakeEventFunc("storage")
	Esubmit                       = domui.MakeEventFunc("submit")
	Esuccess                      = domui.MakeEventFunc("success")
	Esuspend                      = domui.MakeEventFunc("suspend")
	Etimeout                      = domui.MakeEventFunc("timeout")
	Etimeupdate                   = domui.MakeEventFunc("timeupdate")
	Etoggle                       = domui.MakeEventFunc("toggle")
	Etonechange                   = domui.MakeEventFunc("tonechange")
	Etouchcancel                  = domui.MakeEventFunc("touchcancel")
	Etouchend                     = domui.MakeEventFunc("touchend")
	Etouchmove                    = domui.MakeEventFunc("touchmove")
	Etouchstart                   = domui.MakeEventFunc("touchstart")
	Etrack                        = domui.MakeEventFunc("track")
	Etransitioncancel             = domui.MakeEventFunc("transitioncancel")
	Etransitionend                = domui.MakeEventFunc("transitionend")
	Etransitionrun                = domui.MakeEventFunc("transitionrun")
	Etransitionstart              = domui.MakeEventFunc("transitionstart")
	Eunderflow                    = domui.MakeEventFunc("underflow")
	Eunhandledrejection           = domui.MakeEventFunc("unhandledrejection")
	Eunload                       = domui.MakeEventFunc("unload")
	Eunmute                       = domui.MakeEventFunc("unmute")
	Eupgradeneeded                = domui.MakeEventFunc("upgradeneeded")
	Eversionchange                = domui.MakeEventFunc("versionchange")
	Evisibilitychange             = domui.MakeEventFunc("visibilitychange")
	Evoiceschanged                = domui.MakeEventFunc("voiceschanged")
	Evolumechange                 = domui.MakeEventFunc("volumechange")
	Evrdisplayactivate            = domui.MakeEventFunc("vrdisplayactivate")
	Evrdisplayblur                = domui.MakeEventFunc("vrdisplayblur")
	Evrdisplayconnect             = domui.MakeEventFunc("vrdisplayconnect")
	Evrdisplaydeactivate          = domui.MakeEventFunc("vrdisplaydeactivate")
	Evrdisplaydisconnect          = domui.MakeEventFunc("vrdisplaydisconnect")
	Evrdisplayfocus               = domui.MakeEventFunc("vrdisplayfocus")
	Evrdisplaypointerrestricted   = domui.MakeEventFunc("vrdisplaypointerrestricted")
	Evrdisplaypointerunrestricted = domui.MakeEventFunc("vrdisplaypointerunrestricted")
	Evrdisplaypresentchange       = domui.MakeEventFunc("vrdisplaypresentchange")
	Ewaiting                      = domui.MakeEventFunc("waiting")
	Ewebglcontextcreationerror    = domui.MakeEventFunc("webglcontextcreationerror")
	Ewebglcontextlost             = domui.MakeEventFunc("webglcontextlost")
	Ewebglcontextrestored         = domui.MakeEventFunc("webglcontextrestored")
	Ewebkitmouseforcechanged      = domui.MakeEventFunc("webkitmouseforcechanged")
	Ewebkitmouseforcedown         = domui.MakeEventFunc("webkitmouseforcedown")
	Ewebkitmouseforceup           = domui.MakeEventFunc("webkitmouseforceup")
	Ewebkitmouseforcewillbegin    = domui.MakeEventFunc("webkitmouseforcewillbegin")
	Ewheel                        = domui.MakeEventFunc("wheel")
)

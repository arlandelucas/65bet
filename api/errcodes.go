package api

import "errors"

const ( // API error codes
	AECnull = iota

	// authorization

	AEC_auth_absent
	AEC_auth_scheme
	AEC_basic_decode
	AEC_basic_nouser
	AEC_basic_deny
	AEC_token_nouser
	AEC_token_malform
	AEC_token_notsign
	AEC_token_badclaims
	AEC_token_expired
	AEC_token_notyet
	AEC_token_issuer
	AEC_token_error

	// 404

	AEC_nourl

	// 405

	AEC_nomethod

	// GET /signis

	AEC_signis_nobind

	// GET /sendcode

	AEC_sendcode_nobind
	AEC_sendcode_nouser
	AEC_sendcode_update
	AEC_sendcode_code

	// GET /activate

	AEC_activate_nobind
	AEC_activate_nouser
	AEC_activate_oldcode
	AEC_activate_badcode
	AEC_activate_update

	// POST /signup

	AEC_signup_nobind
	AEC_signup_smallsec
	AEC_signup_code
	AEC_signup_sql

	// POST /signin

	AEC_signin_nobind
	AEC_signin_nosecret
	AEC_signin_smallsec
	AEC_signin_nouser
	AEC_signin_activate
	AEC_signin_oldcode
	AEC_signin_badcode
	AEC_signin_denypass
	AEC_signin_sigtime
	AEC_signin_timeout
	AEC_signin_hs256
	AEC_signin_denyhash

	// POST /game/join

	AEC_game_join_nobind
	AEC_game_join_noclub
	AEC_game_join_nouser
	AEC_game_join_noaccess
	AEC_game_join_noalias
	AEC_game_join_sql

	// POST /game/part

	AEC_game_part_nobind
	AEC_game_part_notopened
	AEC_game_part_nouser
	AEC_game_part_noaccess
	AEC_game_part_sql

	// POST /slot/info

	AEC_game_info_nobind
	AEC_game_info_notopened
	AEC_game_info_nouser
	AEC_game_info_noaccess
	AEC_game_info_noprops

	// POST /game/rtp/get

	AEC_game_rdget_nobind
	AEC_game_rdget_notopened
	AEC_game_rdget_noclub
	AEC_game_rdget_nouser
	AEC_game_rdget_noaccess

	// POST /slot/bet/get

	AEC_slot_betget_nobind
	AEC_slot_betget_notopened
	AEC_slot_betget_notslot
	AEC_slot_betget_noaccess

	// POST /slot/bet/set

	AEC_slot_betset_nobind
	AEC_slot_betset_notopened
	AEC_slot_betset_notslot
	AEC_slot_betset_noaccess
	AEC_slot_betset_badbet

	// POST /slot/sel/get

	AEC_slot_selget_nobind
	AEC_slot_selget_notopened
	AEC_slot_selget_notslot
	AEC_slot_selget_noaccess

	// POST /slot/sel/set

	AEC_slot_selset_nobind
	AEC_slot_selset_notopened
	AEC_slot_selset_notslot
	AEC_slot_selset_noaccess
	AEC_slot_selset_badlines

	// POST /slot/mode/set

	AEC_slot_modeset_nobind
	AEC_slot_modeset_notopened
	AEC_slot_modeset_notslot
	AEC_slot_modeset_noaccess
	AEC_slot_modeset_badmode

	// POST /slot/spin

	AEC_slot_spin_nobind
	AEC_slot_spin_notopened
	AEC_slot_spin_notslot
	AEC_slot_spin_noclub
	AEC_slot_spin_nouser
	AEC_slot_spin_noaccess
	AEC_slot_spin_noprops
	AEC_slot_spin_nomoney
	AEC_slot_spin_badbank
	AEC_slot_spin_sqlbank

	// POST /slot/doubleup

	AEC_slot_doubleup_nobind
	AEC_slot_doubleup_notopened
	AEC_slot_doubleup_notslot
	AEC_slot_doubleup_noclub
	AEC_slot_doubleup_nouser
	AEC_slot_doubleup_noaccess
	AEC_slot_doubleup_noprops
	AEC_slot_doubleup_nomoney
	AEC_slot_doubleup_sqlbank

	// POST /slot/collect

	AEC_slot_collect_nobind
	AEC_slot_collect_notopened
	AEC_slot_collect_notslot
	AEC_slot_collect_noaccess
	AEC_slot_collect_denied

	// POST /keno/bet/get

	AEC_keno_betget_nobind
	AEC_keno_betget_notopened
	AEC_keno_betget_notslot
	AEC_keno_betget_noaccess

	// POST /keno/bet/set

	AEC_keno_betset_nobind
	AEC_keno_betset_notopened
	AEC_keno_betset_notslot
	AEC_keno_betset_noaccess
	AEC_keno_betset_badbet

	// POST /keno/sel/get

	AEC_keno_selget_nobind
	AEC_keno_selget_notopened
	AEC_keno_selget_notslot
	AEC_keno_selget_noaccess

	// POST /keno/sel/set

	AEC_keno_selset_nobind
	AEC_keno_selset_notopened
	AEC_keno_selset_notslot
	AEC_keno_selset_noaccess
	AEC_keno_selset_badlines

	// POST /keno/sel/getslice

	AEC_keno_selgetslice_nobind
	AEC_keno_selgetslice_notopened
	AEC_keno_selgetslice_notslot
	AEC_keno_selgetslice_noaccess

	// POST /keno/sel/setslice

	AEC_keno_selsetslice_nobind
	AEC_keno_selsetslice_notopened
	AEC_keno_selsetslice_notslot
	AEC_keno_selsetslice_noaccess
	AEC_keno_selsetslice_badlines

	// POST /keno/spin

	AEC_keno_spin_nobind
	AEC_keno_spin_notopened
	AEC_keno_spin_notslot
	AEC_keno_spin_noclub
	AEC_keno_spin_nouser
	AEC_keno_spin_noaccess
	AEC_keno_spin_noprops
	AEC_keno_spin_nomoney
	AEC_keno_spin_badbank
	AEC_keno_spin_sqlbank

	// POST /prop/get

	AEC_prop_get_nobind
	AEC_prop_get_noclub
	AEC_prop_get_nouser
	AEC_prop_get_noaccess

	// POST /prop/wallet/get

	AEC_prop_walletget_nobind
	AEC_prop_walletget_noclub
	AEC_prop_walletget_nouser
	AEC_prop_walletget_noaccess

	// POST /prop/wallet/add

	AEC_prop_walletadd_nobind
	AEC_prop_walletadd_limit
	AEC_prop_walletadd_noclub
	AEC_prop_walletadd_nouser
	AEC_prop_walletadd_noaccess
	AEC_prop_walletadd_noprops
	AEC_prop_walletadd_nomoney
	AEC_prop_walletadd_sql

	// POST /prop/al/get

	AEC_prop_alget_nobind
	AEC_prop_alget_noclub
	AEC_prop_alget_nouser
	AEC_prop_alget_noaccess

	// POST /prop/al/set

	AEC_prop_alset_nobind
	AEC_prop_alset_noclub
	AEC_prop_alset_nouser
	AEC_prop_alset_noaccess
	AEC_prop_alset_noprops
	AEC_prop_alset_nolevel
	AEC_prop_alset_sql

	// POST /prop/rtp/get

	AEC_prop_rtpget_nobind
	AEC_prop_rtpget_noclub
	AEC_prop_rtpget_nouser
	AEC_prop_rtpget_noaccess

	// POST /prop/rtp/set

	AEC_prop_rtpset_nobind
	AEC_prop_rtpset_noclub
	AEC_prop_rtpset_nouser
	AEC_prop_rtpset_noaccess
	AEC_prop_rtpset_noprops
	AEC_prop_rtpset_sql

	// POST /user/rename

	AEC_user_rename_nobind
	AEC_user_rename_nouser
	AEC_user_rename_noaccess
	AEC_user_rename_update

	// POST /user/secret

	AEC_user_secret_nobind
	AEC_user_secret_smallsec
	AEC_user_secret_nouser
	AEC_user_secret_noaccess
	AEC_user_secret_notequ
	AEC_user_secret_update

	// POST /user/delete

	AEC_user_delete_nobind
	AEC_user_delete_nouser
	AEC_user_delete_noaccess
	AEC_user_delete_nosecret
	AEC_user_delete_sqluser
	AEC_user_delete_sqllock
	AEC_user_delete_sqlprops
	AEC_user_delete_sqlgames

	// POST /club/is

	AEC_club_is_nobind

	// POST /club/info

	AEC_club_info_nobind
	AEC_club_info_noclub
	AEC_club_info_noaccess

	// POST /club/rename

	AEC_club_rename_nobind
	AEC_club_rename_noclub
	AEC_club_rename_noaccess
	AEC_club_rename_update

	// POST /club/cashin

	AEC_club_cashin_nobind
	AEC_club_cashin_nosum
	AEC_club_cashin_noclub
	AEC_club_cashin_noaccess
	AEC_club_cashin_bankout
	AEC_club_cashin_fundout
	AEC_club_cashin_lockout
	AEC_club_cashin_sqlbank
	AEC_club_cashin_sqllog
)

var (
	Err404       = errors.New("page not found")
	Err405       = errors.New("method not allowed")
	ErrNoClub    = errors.New("club with given ID does not found")
	ErrNoUser    = errors.New("user with given ID does not found")
	ErrNoProps   = errors.New("properties for given user and club does not found")
	ErrNoAddSum  = errors.New("no sum to change balance of bank or fund or deposit")
	ErrNoMoney   = errors.New("not enough money on balance")
	ErrBankOut   = errors.New("not enough money at bank")
	ErrFundOut   = errors.New("not enough money at jackpot fund")
	ErrLockOut   = errors.New("not enough money at deposit")
	ErrNotSlot   = errors.New("specified GID refers to non-slot game")
	ErrNotKeno   = errors.New("specified GID refers to non-keno game")
	ErrNoAccess  = errors.New("no access rights for this feature")
	ErrNoLevel   = errors.New("admin have no privilege to modify specified access level to user")
	ErrNotConf   = errors.New("password confirmation does not pass")
	ErrZero      = errors.New("given value is zero")
	ErrTooBig    = errors.New("given value exceeds the limit")
	ErrNoAliase  = errors.New("no game alias")
	ErrNotOpened = errors.New("game with given ID is not opened")
	ErrBadBank   = errors.New("can not generate spin with current bank balance")
)
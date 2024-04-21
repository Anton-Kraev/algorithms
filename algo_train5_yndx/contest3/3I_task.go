package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	gamesCnt, totalGoals, scoreOpens int
}

type Player struct {
	team       string
	scoreOpens int
	goals      []int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var result []string
	stats := Stats{make(map[string]*Team), make(map[string]*Player)}
	var gameInfo []string
	gameInfoLines := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if gameInfoLines > 0 {
			gameInfo = append(gameInfo, line)
		}
		if line[0] == '"' {
			gameInfo = []string{line}
			splitted := strings.Split(line, " ")
			var score1, score2 int
			fmt.Sscanf(splitted[len(splitted)-1], "%d:%d", &score1, &score2)
			gameInfoLines = score1 + score2 + 1
		}
		gameInfoLines--
		if gameInfoLines == 0 {
			stats.Update(gameInfo)
		} else if gameInfoLines < 0 {
			result = append(result, stats.GetInfo(line))
		}
	}

	for _, s := range result {
		fmt.Println(s)
	}
}

type Stats struct {
	teams   map[string]*Team
	players map[string]*Player
}

func (stats Stats) Update(gameInfo []string) {
	var team1, team2 string
	var score1, score2 int
	mainInfo := strings.Split(gameInfo[0], "\"")
	team1, team2 = mainInfo[1], mainInfo[3]
	fmt.Sscanf(strings.TrimSpace(mainInfo[4]), "%d:%d", &score1, &score2)
	_, team1InStats := stats.teams[team1]
	if !team1InStats {
		stats.teams[team1] = &Team{0, 0, 0}
	}
	stats.teams[team1].gamesCnt++
	stats.teams[team1].totalGoals += score1
	_, team2InStats := stats.teams[team2]
	if !team2InStats {
		stats.teams[team2] = &Team{0, 0, 0}
	}
	stats.teams[team2].gamesCnt++
	stats.teams[team2].totalGoals += score2

	openMinute := 91
	openPlayer := Player{"", 0, []int{}}
	openPlayerName := ""
	for i, goal := range gameInfo[1:] {
		goalInfo := strings.Split(goal, " ")
		player := strings.Join(goalInfo[:len(goalInfo)-1], " ")
		minute, _ := strconv.Atoi(strings.ReplaceAll(goalInfo[len(goalInfo)-1], "'", ""))
		_, inStats := stats.players[player]
		if !inStats {
			if i < score1 {
				stats.players[player] = &Player{team1, 0, []int{}}
			} else {
				stats.players[player] = &Player{team2, 0, []int{}}
			}
		}
		stats.players[player].goals = append(stats.players[player].goals, minute)
		if minute < openMinute {
			openMinute = minute
			openPlayer = *stats.players[player]
			openPlayerName = player
		}
	}

	if openPlayerName != "" {
		stats.players[openPlayerName].scoreOpens++
		stats.teams[openPlayer.team].scoreOpens++
	}
}

func (stats Stats) GetInfo(query string) string {
	tokens := strings.Split(query, " ")
	switch tokens[0] {
	case "Total":
		name := strings.ReplaceAll(strings.Join(tokens[3:], " "), "\"", "")
		switch tokens[2] {
		case "for":
			_, isTeam := stats.teams[name]
			if !isTeam {
				return "0"
			}
			return stats.totalTeamGoals(name)
		case "by":
			_, isPlayer := stats.players[name]
			if !isPlayer {
				return "0"
			}
			return stats.totalPlayerGoals(name)
		}
	case "Mean":
		name := strings.ReplaceAll(strings.Join(tokens[5:], " "), "\"", "")
		switch tokens[4] {
		case "for":
			_, isTeam := stats.teams[name]
			if !isTeam {
				return "0"
			}
			return stats.meanTeamGoals(name)
		case "by":
			_, isPlayer := stats.players[name]
			if !isPlayer {
				return "0"
			}
			return stats.meanPlayerGoals(name)
		}
	case "Goals":
		switch tokens[2] {
		case "minute":
			name := strings.Join(tokens[5:], " ")
			_, isPlayer := stats.players[name]
			if !isPlayer {
				return "0"
			}
			min, _ := strconv.Atoi(tokens[3])
			return stats.goalsOnMinute(min, name)
		case "first":
			name := strings.Join(tokens[6:], " ")
			_, isPlayer := stats.players[name]
			if !isPlayer {
				return "0"
			}
			min, _ := strconv.Atoi(tokens[3])
			return stats.firstMinutesGoals(min, name)
		case "last":
			name := strings.Join(tokens[6:], " ")
			_, isPlayer := stats.players[name]
			if !isPlayer {
				return "0"
			}
			min, _ := strconv.Atoi(tokens[3])
			return stats.lastMinutesGoals(min, name)
		}
	case "Score":
		name := strings.ReplaceAll(strings.Join(tokens[3:], " "), "\"", "")
		_, isTeam := stats.teams[name]
		_, isPlayer := stats.players[name]
		if isTeam && tokens[3][0] == '"' {
			return stats.teamScoreOpens(name)
		} else if isPlayer && tokens[3][0] != '"' {
			return stats.playerScoreOpens(name)
		}
		return "0"
	}
	return ""
}

func (stats Stats) totalTeamGoals(team string) string {
	return strconv.Itoa(stats.teams[team].totalGoals)
}

func (stats Stats) meanTeamGoals(team string) string {
	return strconv.FormatFloat(
		float64(stats.teams[team].totalGoals)/float64(stats.teams[team].gamesCnt),
		'f', 4, 64,
	)
}

func (stats Stats) totalPlayerGoals(player string) string {
	return strconv.Itoa(len(stats.players[player].goals))
}

func (stats Stats) meanPlayerGoals(player string) string {
	return strconv.FormatFloat(
		float64(len(stats.players[player].goals))/float64(stats.teams[stats.players[player].team].gamesCnt),
		'f', 4, 64,
	)
}

func (stats Stats) goalsOnMinute(min int, player string) string {
	return stats.goalsInTimeInterval(min, min, player)
}

func (stats Stats) firstMinutesGoals(min int, player string) string {
	return stats.goalsInTimeInterval(1, min, player)
}

func (stats Stats) lastMinutesGoals(min int, player string) string {
	return stats.goalsInTimeInterval(91-min, 90, player)
}

func (stats Stats) goalsInTimeInterval(fromMin, toMin int, player string) string {
	cnt := 0
	for _, goal := range stats.players[player].goals {
		if goal >= fromMin && goal <= toMin {
			cnt++
		}
	}
	return strconv.Itoa(cnt)
}

func (stats Stats) teamScoreOpens(team string) string {
	return strconv.Itoa(stats.teams[team].scoreOpens)
}

func (stats Stats) playerScoreOpens(player string) string {
	return strconv.Itoa(stats.players[player].scoreOpens)
}

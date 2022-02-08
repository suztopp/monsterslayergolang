# PLANNING THE APP

## MONSTER SLAYER APP

FLOW

### Start Game

Player and Monster Participants

Start New Round

Ask Player to Attack, Heal, Special Attack

Update Monster health after each choice or attack - calculated randomly

Healing will also adjust random number for player health update

Monster will then attack player after player choice - calculate damage with random number

Update Player health each time this happens

Check for a winner after the game - Player or Monster based on health

If both health amounts are 0 or below - monster wins

No winner? Choose action

End Game? Output summary and write log file

---

### Modules and Packages to Create

MODULES

url/monsterslayer

PACKAGES

/main

/interaction - userinput and output

/actions - calculating health etc and updated hit points

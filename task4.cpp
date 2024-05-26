#include <iostream>

using namespace std;

bool unpredictable(int round_number, bool self_choices[], bool enemy_choices[]) {
    return rand() % 2;
}

bool reciprocity(int round_number, bool self_choices[], bool enemy_choices[]) {
    if (!enemy_choices[round_number - 2]) {
        return false;
    }
    return true;
}

bool mastery(int round_number, bool self_choices[], bool enemy_choices[]) {
    if (round_number == 1) {
        return true;
    }

    int self_betrayals = 0, enemy_betrayals = 0;
    for(int i = 0; i < round_number; i++) {
        if (!self_choices[i]) {
            self_betrayals++;
        }
        if (!enemy_choices[i]) {
            enemy_betrayals++;
        }
    }

    if (self_betrayals > enemy_betrayals) {
        return false;
    }

    return true;
}

void game(bool (*tact1)(int, bool[], bool[]), bool (*tact2)(int, bool[], bool[]), int& pl1_score, int& pl2_score) {
    int rounds = rand() % 101 + 100;

    bool pl1_move[101], pl2_move[101];

    for (int i = 1; i <= rounds; i++) {
        bool decision1 = tact1(i, pl1_move, pl2_move);
        bool decision2 = tact2(i, pl2_move, pl1_move);

        pl1_move[i - 1] = decision1;
        pl2_move[i - 1] = decision2;

        if (decision1 && decision2) {
            pl1_score += 24;
            pl2_score += 24;
        }
        if (decision1 && !decision2) {
            pl1_score += 0;
            pl2_score += 20;
        }
        if (!decision1 && decision2) {
            pl1_score += 20;
            pl2_score += 0;
        } 
        else {
            pl1_score += 4;
            pl2_score += 4;
        }
    }
}

int main() {
    srand(time(0));
    
    int unpred_score = 0, recip_score = 0, mast_score = 0;

    game(unpredictable, reciprocity, unpred_score, recip_score);
    game(unpredictable, mastery, unpred_score, mast_score);
    game(reciprocity, mastery, recip_score, mast_score);

    cout << "Счёт игрока с тактикой unpredictable: " << unpred_score << endl;
    cout << "Счёт игрока с тактикой reciprocity: " << recip_score << endl;
    cout << "Счёт игрока с тактикой mastery: " << mast_score << endl;
}
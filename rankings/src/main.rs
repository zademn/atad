use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */

fn climbingLeaderboard(ranked: &[i32], player: &[i32]) -> Vec<i32> {
    let n_players = ranked.len();
    let n_games = player.len();

    let mut player_ranks = vec![0; n_players];
    player_ranks[0] = 1;
    for p1 in 1..n_players {
        if ranked[p1] == ranked[p1 - 1] {
            player_ranks[p1] = player_ranks[p1 - 1];
            continue;
        }
        if ranked[p1] < ranked[p1 - 1] {
            player_ranks[p1] = player_ranks[p1 - 1] + 1;
            continue;
        }
    }
    println!("{:?}", player_ranks);
    let mut current_ranks = vec![1; n_games];
    for (j, game) in player.iter().enumerate() {
        for (i, p_score) in ranked.iter().enumerate().rev() {
            if game < p_score {
                current_ranks[j] = player_ranks[i] + 1;
                break;
            }
        }
    }
    println!("{:?}", current_ranks);
    current_ranks
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    // let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let _ranked_count = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<i32>()
        .unwrap();

    let ranked: Vec<i32> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    let _player_count = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim()
        .parse::<i32>()
        .unwrap();

    let player: Vec<i32> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .trim_end()
        .split(' ')
        .map(|s| s.to_string().parse::<i32>().unwrap())
        .collect();

    let result = climbingLeaderboard(&ranked, &player);

    // for i in 0..result.len() {
    //     write!(&mut fptr, "{}", result[i]).ok();

    //     if i != result.len() - 1 {
    //         writeln!(&mut fptr).ok();
    //     }
    // }

    // writeln!(&mut fptr).ok();
}

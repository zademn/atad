use std::io::{self, BufRead};

/*
 * Complete the 'surfaceArea' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY A as parameter.
 */

fn surfaceArea(A: &[Vec<i32>]) -> i32 {
    let mut s = 0;
    let n_rows = A.len();
    let n_cols = A[0].len();
    for i in 0..n_rows {
        for j in 0..n_cols {
            // add top and bot
            if A[i][j] != 0 {
                s += 2;
            }
            // add outer sides
            if i == 0 {
                s += A[i][j];
            }
            if i == n_rows - 1 {
                s += A[i][j];
            }
            if j == 0 {
                s += A[i][j];
            }
            if j == n_cols - 1 {
                s += A[i][j];
            }
            // add inner sides
            if i > 0 {
                let v = A[i][j] - A[i - 1][j];
                if v > 0 {
                    s += v;
                }
            }
            if i < n_rows - 1 {
                let v = A[i][j] - A[i + 1][j];
                if v > 0 {
                    s += v;
                }
            }
            if j > 0 {
                let v = A[i][j] - A[i][j - 1];
                if v > 0 {
                    s += v;
                }
            }
            if j < n_cols - 1 {
                let v = A[i][j] - A[i][j + 1];
                if v > 0 {
                    s += v;
                }
            }
        }
    }
    println!("{}", s);
    s
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    // let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let first_multiple_input: Vec<String> = stdin_iterator
        .next()
        .unwrap()
        .unwrap()
        .split(' ')
        .map(|s| s.to_string())
        .collect();

    let H = first_multiple_input[0].trim().parse::<i32>().unwrap();

    let W = first_multiple_input[1].trim().parse::<i32>().unwrap();

    let mut A: Vec<Vec<i32>> = Vec::with_capacity(H as usize);

    for i in 0..H as usize {
        A.push(Vec::with_capacity(W as usize));

        A[i] = stdin_iterator
            .next()
            .unwrap()
            .unwrap()
            .trim_end()
            .split(' ')
            .map(|s| s.to_string().parse::<i32>().unwrap())
            .collect();
    }

    let _result = surfaceArea(&A);

    // writeln!(&mut fptr, "{}", result).ok();
}

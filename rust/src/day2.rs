pub fn gravity_assist() {
    let input = vec![
        1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,6,19,23,2,23,6,27,1,5,27,31,1,10,31,35,2,6,35,39,1,
        39,13,43,1,43,9,47,2,47,10,51,1,5,51,55,1,55,10,59,2,59,6,63,2,6,63,67,1,5,67,71,2,9,71,75,1,75,
        6,79,1,6,79,83,2,83,9,87,2,87,13,91,1,10,91,95,1,95,13,99,2,13,99,103,1,103,10,107,2,107,10,111,1,
        111,9,115,1,115,2,119,1,9,119,0,99,2,0,14,0
    ];

    let sequence = inputs(input.clone(), 12, 2);
    let resp = intcode(sequence)[0];
    println!("final output: {}", resp);

    let resp2 = find_inputs(input, 19690720);
    println!("inputs: {:?}, code: {}", resp2, (100 * resp2.0 + resp2.1));
}

fn find_inputs(input: Vec<i32>, target: i32) -> (i32, i32) {
    for noun in 0..99 {
        for verb in 0..99 {
            let s = intcode(inputs(input.clone(), noun, verb));
            if s[0] == target {
                return (noun, verb)
            }
        }
    }
    return (-1, -1)
}

fn inputs(mut input: Vec<i32>, noun: i32, verb: i32) -> Vec<i32> {
    input[1] = noun;
    input[2] = verb;
    return input;
}

fn intcode(input: Vec<i32>) -> Vec<i32> {
    return rec(input, 0)
}

fn rec(mut input: Vec<i32>, offset: usize) -> Vec<i32> {
    if input[offset] == 99 {
        return input
    }
    let a = input[offset + 1] as usize;
    let b = input[offset + 2] as usize;
    let c = input[offset + 3] as usize;
    if input[offset] == 1 {
        input[c] = input[a] + input[b];
    } else if input[offset] == 2 {
        input[c] = input[a] * input[b];
    } else {
        panic!("unknown operation")
    }
    return rec(input, offset + 4);
}

#[cfg(test)]
#[test]
fn advent002_1() {
    assert_eq!(intcode(vec![1,0,0,0,99]), vec![2,0,0,0,99]);
    assert_eq!(intcode(vec![2,3,0,3,99]), vec![2,3,0,6,99]);
    assert_eq!(intcode(vec![2,4,4,5,99,0]), vec![2,4,4,5,99,9801]);
    assert_eq!(intcode(vec![1,1,1,4,99,5,6,0,99]), vec![30,1,1,4,2,5,6,0,99]);
}

use std::{error::Error, fs};

static INPUT_PATHS: &[&str] = &[
    "level1/level1_0_example",
    "level1/level1_1_small",
    "level1/level1_2_large",
];

struct Input {
    sequence_count: u32,
    sequences: Vec<Vec<u32>>,
}

struct Output {
    total_times: Vec<u32>,
}

pub fn run() -> Result<(), Box<dyn Error>> {
    for path in INPUT_PATHS {
        let input = read_input(path)?;
        let output = compute_total_times(&input);
        write_output(path, &output)?;
    }
    Ok(())
}

fn read_input(path: &str) -> Result<Input, Box<dyn Error>> {
    let input_str = fs::read_to_string(path.to_owned() + ".in")?;
    let mut lines = input_str.lines();

    let sequence_count = lines.next().ok_or("malformed")?.parse::<u32>()?;
    let mut sequences = Vec::with_capacity(sequence_count as usize);

    for line in lines {
        let r: Result<Vec<u32>, _> = line.split_ascii_whitespace().map(str::parse).collect();
        sequences.push(r?);
    }

    Ok(Input {
        sequence_count,
        sequences,
    })
}

fn compute_total_times(input: &Input) -> Output {
    let mut total_times = Vec::with_capacity(input.sequence_count as usize);

    for seq in &input.sequences {
        total_times.push(seq.iter().sum());
    }

    Output { total_times }
}

fn write_output(path: &str, output: &Output) -> Result<(), Box<dyn Error>> {
    fs::write(
        path.to_owned() + ".out",
        output
            .total_times
            .iter()
            .map(u32::to_string)
            .collect::<Vec<_>>()
            .join("\n"),
    )?;
    Ok(())
}

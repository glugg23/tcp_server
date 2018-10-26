use std::net::TcpListener;
use std::io::{BufReader, BufRead};

fn main() {
    let listener = TcpListener::bind("127.0.0.1:9999").unwrap();

    for stream in listener.incoming() {
        let stream = stream.expect("Unable to accept");
        
        let mut reader = BufReader::new(&stream);
        let mut response = String::new();
        reader.read_line(&mut response).expect("Read error");

        println!("{}", response);
    }
}
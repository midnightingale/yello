//import {background, createCanvas, loadImage, windowHeight, windowWidth} from "p5/global";

let levelImg;
let pellets = [];
let entities = [];

const levelStr = "000 000 000 000 000 000 000 000 000 100 021 100 000 000 000 000 000 000 000 000 000\n" +
    "000 100 100 100 100 100 100 100 100 100 000 100 100 100 100 100 100 100 100 100 000\n" +
    "000 100 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 100 000\n" +
    "000 100 003 100 002 100 002 100 100 100 002 100 100 100 002 100 002 100 003 100 000\n" +
    "000 100 002 002 002 100 002 002 002 100 002 100 002 002 002 100 002 002 002 100 000\n" +
    "000 100 002 100 100 100 002 100 002 100 002 100 002 100 002 100 100 100 002 100 000\n" +
    "000 100 002 002 002 002 002 100 002 002 002 002 002 100 002 002 002 002 002 100 000\n" +
    "000 100 002 100 002 100 100 100 100 100 002 100 100 100 100 100 002 100 002 100 000\n" +
    "000 100 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 100 000\n" +
    "000 100 100 100 002 100 002 100 100 100 100 100 100 100 002 100 002 100 100 100 000\n" +
    "000 000 000 100 002 100 002 002 000 000 010 000 000 002 002 100 002 100 000 000 000\n" +
    "100 100 100 100 002 100 100 002 100 100 001 100 100 002 100 100 002 100 100 100 100\n" +
    "020 000 000 000 002 000 000 002 100 011 012 013 100 002 000 000 002 000 000 000 020\n" +
    "100 100 100 100 002 100 100 002 100 100 100 100 100 002 100 100 002 100 100 100 100\n" +
    "000 100 002 002 002 000 000 002 000 000 022 000 000 002 000 000 002 002 002 100 000\n" +
    "000 100 002 100 002 100 100 002 100 100 100 100 100 002 100 100 002 100 002 100 000 \n" +
    "000 100 002 002 002 100 002 002 000 000 004 000 000 002 002 100 002 002 002 100 000\n" +
    "000 100 100 100 002 100 002 100 100 100 100 100 100 100 002 100 002 100 100 100 000 \n" +
    "000 000 100 002 002 002 002 002 002 002 100 002 002 002 002 002 002 002 100 000 000 \n" +
    "000 000 100 002 100 100 002 100 100 002 100 002 100 100 002 100 100 002 100 000 000\n" +
    "100 100 100 002 002 002 002 002 100 002 100 002 100 002 002 002 002 002 100 100 100\n" +
    "100 002 002 002 100 100 100 002 100 002 100 002 100 002 100 100 100 002 002 002 100\n" +
    "100 003 100 002 002 002 002 002 002 002 002 002 002 002 002 002 002 002 100 003 100\n" +
    "100 002 002 002 100 100 100 100 100 100 000 100 100 100 100 100 100 002 002 002 100\n" +
    "100 100 100 100 100 000 000 000 000 100 021 100 000 000 000 000 100 100 100 100 100"

let level = [];
let levelHeight, levelWidth;
let block_size;


function preload() {
    levelImg = loadImage('images/level.png')
    level = levelStr.split('\n').map(str => str.split(' '))
    print(level)
}

function calc_block_size() {
    levelHeight = windowHeight / 2
    levelWidth = windowHeight / 2 / levelImg.height * levelImg.width
    block_size = levelHeight / 25
}

function setup() {
    calc_block_size();
    createCanvas(levelWidth, levelHeight);
}

function windowResized() {
    calc_block_size();
    createCanvas(levelWidth, levelHeight);
}


function draw() {
    fill(255, 204, 0)
    for (let i = 0; i < 25; i++) {
        for (let j = 0; j < 21; j++) {
            if (level[i][j] == '100') {
                rect(j * block_size, i * block_size, block_size, block_size)
            }
        }
    }

    // background(levelImg);
    // x++
}

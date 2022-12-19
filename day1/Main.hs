module Main where

main :: IO ()
main = do
    raw <- readFile "input.text"
    let currentCalorieSum = 0
    let maxCalories = 0
    input <- getLine
    let nums = map read $ lines raw :: [Int]
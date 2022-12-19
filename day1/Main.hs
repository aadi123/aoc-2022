module Main where

main :: IO ()
main = do
    let currentCalorieSum = 0
    let maxCalories = 0
    raw <- readFile "input.txt"
    putStrLn raw
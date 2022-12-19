module Main where

main :: IO ()
main = do
    raw <- readFile "input.txt"
    let currentCalorieSum = 0
    let maxCalories = 0
    putStrLn raw
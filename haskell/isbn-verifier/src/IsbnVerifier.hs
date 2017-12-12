module IsbnVerifier (isbn) where

import Data.Char (digitToInt)

isbn :: String -> Bool
isbn s = validDigits s && isbnValue s `mod` 11 == 0

validDigits :: String -> Bool
validDigits s
  | 'X' `elem` take (length s - 1) s = False
  | otherwise = True

isbnValue :: String -> Integer
isbnValue = calculate . map change . keep "0123456789X"

mult :: (Integer, Integer) -> Integer
mult (i, n) = i * n

calculate :: [Integer] -> Integer
calculate series = sum $ zipWith (curry mult) [10, 9 .. 0] series

keep :: String -> String -> String
keep characters target = [ char | char <- target, char `elem` characters]

change :: Char -> Integer
change 'X' = 10
change c = toInteger $ digitToInt c
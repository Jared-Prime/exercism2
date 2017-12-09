module IsbnVerifier (isbn) where

isbn :: String -> Bool
isbn s = validDigits s && (isbnValue s) `mod` 11 == 0

validDigits :: String -> Bool
validDigits s
  | any (\digit -> digit == 'X') (take (length(s) - 1) s) = False
  | otherwise = True

isbnValue :: String -> Integer
isbnValue = calculate . map change . keep "0123456789X"

mult :: (Integer, Integer) -> Integer
mult (i, n) = i * n

calculate :: [Integer] -> Integer
calculate series = sum $ map mult (zip [10, 9..0] series)

keep :: String -> String -> String
keep characters target = [ char | char <- target, char `elem` characters]

change :: Char -> Integer
change '0' = 0
change '1' = 1
change '2' = 2
change '3' = 3
change '4' = 4
change '5' = 5
change '6' = 6
change '7' = 7
change '8' = 8
change '9' = 9
change 'X' = 10
change _ = 0
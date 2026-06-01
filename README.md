TrimmedMean function for MSDS 431 week 9 assignment

Inputs:  
    number slice T (accepts various numerical variable types)  
    lowTrim (percent of lowest values in T to be trimmed)  
    highTrim (percent of highest values in T to be trimmed. If no value provided, will use the same value as lowTrim)  

outputs:  
    float64 value containing the mean of slice T after trimming lowTrim% of the lowest values and highTrim% of the highest values in T

trimmedmean.TrimmedMean(T, lowTrim, highTrim)

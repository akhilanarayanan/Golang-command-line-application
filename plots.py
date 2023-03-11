import sys
import pandas as pd
from matplotlib import pyplot as plt
import numpy as np

def main():
  # df = pd.read_csv("small_output.csv")
  df = pd.read_csv("large_output.csv")
  # Average size of bytes across all the entries
  print("average byte size:", df.iloc[:, 4].mean())
  
  # Create a line plot of each request and the time taken to get the response
  durationCol = (df.iloc[:, 3] - df.iloc[:, 2])
  plt.plot(durationCol / (10**6))  # converts nanoseconds to milliseconds
  plt.title('Request Response Times')
  plt.ylabel('Time (ms)')
  plt.show()
  
  # compute the CDF of the response time column
  x, y = cdf(durationCol / (10**6))  # converts nanoseconds to milliseconds

  # find the specified percentiles
  percentiles = [50, 70, 75, 80, 85, 90, 95, 99]
  for p in percentiles:
    np.percentile(durationCol, p)

  # plot the CDF
  plt.plot(x, y)
  plt.title('CDF of Response Times')
  plt.xlabel('Time (ms)')
  plt.ylabel('Percentile')
  plt.show()  
  
def cdf(data):
  n = len(data)
  x = np.sort(data)
  y = np.arange(1, n + 1) / n
  return x, y
  
if __name__ == '__main__':
  main()

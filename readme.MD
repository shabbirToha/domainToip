# Domain to IP Resolver

This Go script reads domain names from an input file (`domains.txt`), resolves them to their respective IP addresses, and saves the results in an output file (`output.txt`). It includes a retry mechanism to handle intermittent DNS resolution failures.

## Features

- Reads domain names from `domains.txt`
- Resolves each domain to an IP address
- Retries up to 3 times with a 2-second delay if DNS resolution fails
- Saves results in `output.txt`

## Requirements

- Go (1.18 or later)

## Installation

1. Clone this repository or download the script.
2. Ensure you have Go installed.

## Usage

1. Create a `domains.txt` file with one domain per line, e.g.:
   ```txt
   google.com
   example.com
   github.com
   ```
2. Run the script:
   ```sh
   go run main.go
   ```
3. Check the output file `output.txt` for results.

## Example Output

```txt
google.com: 142.250.64.78
example.com: 93.184.216.34
github.com: 140.82.121.4
```

## Notes

- If a domain cannot be resolved, it will be marked as `Failed to resolve` in the output.
- You can modify the retry attempts and delay in the `resolveIPWithRetry` function.



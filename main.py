# reducing here is just running a sum over digits repeatedly until we have a
# single-digit number

def reduceNumber(num):
    while True:
        s = 0
        while True:
            s = s + num % 10
            num //= 10
            if num == 0:
                break
        num = s
        if num < 10:
            break
    return num

def reduceNumber2(num):
    def sumDigits(n):
        s = 0
        while True:
            s += n % 10
            n //= 10
            if n == 0:
                return s

    while True:
        while True:
            num = sumDigits(num)
            if num < 10:
                return num

if __name__ == '__main__':
    try:
        while True:
            sln = [reduceNumber, reduceNumber2]
            input_num = int(input('int << '))
            for idx, fn in enumerate(sln):
                            print(f'Answer #{idx+1}', fn(input_num), sep=' ')
            print()
    except KeyboardInterrupt:
        exit()

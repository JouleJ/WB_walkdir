import os

def generate_test(path, n):
    os.makedirs(path)
    if n > 0:
        generate_test(os.path.join(path, '0'), n - 1)
        generate_test(os.path.join(path, '1'), n - 1)

generate_test('test', 20)

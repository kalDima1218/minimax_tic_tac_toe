def get_score(board, depth):
    """
    This function calculates the score of the current board state.
    :param board: current state of the board
    :param depth: current depth in the game tree
    :return: score of the current board state
    """
    if check_win(board):
        return 10 - depth
    elif check_win(board, 'O'):
        return depth - 10
    else:
        return 0

def make_move(board, move, player):
    """
    This function applies a move to the board and returns the new board.
    :param board: current state of the board
    :param move: tuple representing the move to be made
    :param player: symbol representing the player making the move
    :return: new board with the move applied
    """
    new_board = [row[:] for row in board]
    new_board[move[0]][move[1]] = player
    return new_board

def get_possible_moves(board):
    """
    This function returns a list of all possible moves that can be made on the board.
    :param board: current state of the board
    :return: list of possible moves
    """
    moves = []
    for i in range(3):
        for j in range(3):
            if board[i][j] == ' ':
                moves.append((i, j))
    return moves

def check_win(board):
    """
    This function checks if either player has won the game.
    :param board: current state of the board
    :return: boolean indicating whether either player has won
    """
    # check rows
    for row in board:
        if row.count('X') == 3 or row.count('O') == 3:
            return True

    # check columns
    for i in range(3):
        if board[0][i] == board[1][i] == board[2][i] and board[0][i] != ' ':
            return True

    # check diagonals
    if board[0][0] == board[1][1] == board[2][2] and board[0][0] != ' ':
        return True
    if board[0][2] == board[1][1] == board[2][0] and board[0][2] != ' ':
        return True

    return False

def min_max(board, depth, is_maximizing):
    """
    This function implements the Minimax algorithm for Tic Tac Toe game.
    :param board: current state of the board
    :param depth: current depth in the game tree
    :param is_maximizing: boolean indicating whether it's the maximizing player's turn or not
    :return: tuple containing the best score and corresponding best move
    """
    # base case: check if the game is over or maximum depth is reached
    if check_win(board):
        return get_score(board, depth), None
    elif depth == 0:
        return 0, None

    # recursive case: generate all possible moves and evaluate them
    if is_maximizing:
        best_score = float('-inf')
        for move in get_possible_moves(board):
            new_board = make_move(board, move, 'X')
            score, _ = min_max(new_board, depth - 1, False)
            if score > best_score:
                best_score, best_move = score, move
    else:
        best_score = float('inf')
        for move in get_possible_moves(board):
            new_board = make_move(board, move, 'O')
            score, _ = min_max(new_board, depth - 1, True)
            if score < best_score:
                best_score, best_move = score, move

    return best_score, best_move

board = [
    [' ', ' ', ' '],
    [' ', ' ', ' '],
    [' ', ' ', ' ']
]

best_score, best_move = min_max(board, 5, True)

print(f"Best score: {best_score}")
print(f"Best move: {best_move}")

export async function load({ params }) {
    const number = params.number;
    
    // console.log(`Server load for simulation ${number}`);
    
    return {
        data: {
            rounds: [
                {
                    attack: {
                        moves: [
                            {
                                from: {
                                    row: 0,
                                    col: 0,
                                    value: 1,
                                },
                                to: {
                                    row: 1,
                                    col: 0,
                                    value: 1,
                                },
                                is_winner: true,
                                statistics: {
                                    scores: {
                                        "Nikitos": 2,
                                        "Matvey": 1,
                                    }
                                }
                            }
                        ]
                    },
                    defence: {
                        cells: [
                            {
                                row: 0, 
                                col: 0, 
                                value: 3
                            }
                        ],
                        statistics: {
                            scores: {
                                "Nikitos": 2,
                                "Matvey": 1,
                            }
                        }
                    }
                },
                {
                    attack: {
                        moves: [
                            {
                                from: {
                                    row: 2,
                                    col: 2,
                                    value: 1,
                                },
                                to: {
                                    row: 1,
                                    col: 1,
                                    value: 1,
                                },
                                is_winner: true,
                                statistics: {
                                    scores: {
                                        "Nikitos": 2,
                                        "Matvey": 2,
                                    }
                                }
                            }
                        ]   
                    },
                    defence: {
                        cells: [
                            {
                                row: 2, 
                                col: 2, 
                                value: 3
                            }
                        ],
                        statistics: {
                            scores: {
                                "Nikitos": 2,
                                "Matvey": 2,
                            }
                        }
                    }
                }
            ],
            map: {
                name: "small_map_3x3",
                ncols: 10,
                nrows: 20,
                fieldMask: [
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,

                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,

                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,

                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                    true, true, true, true, true, true, true, true, true, true,
                ],
                bigCells: [
                    {row: 0, col: 0},
                    {row: 2, col: 2},
                    {row: 19, col: 9},
                ],
            },
            start_positions: [
                {row: 0, col: 0},
                {row: 2, col: 2},
            ],
            names: [
                "Nikitos",
                "Matvey",
            ]
        }
    };
}
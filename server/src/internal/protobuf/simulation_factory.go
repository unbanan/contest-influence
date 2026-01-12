package protobuf

import (
    "encoding/json"
    "fmt"
    "strings"
)

type SimulationFactory struct{}

type SimulationData struct {
    Players   []string `json:"players"`
    MaxTurns  int      `json:"max_turns"`
    CreatedAt string   `json:"created_at"`
}

func (f *SimulationFactory) GenerateAllSimulations() string {
    mapIDs := []string{
        "cc10a199-5046-45f6-bedf-ad037f0b0ee8",
        "5d9d7950-2209-4d8f-af33-06f8b3e2ed0a", 
        "fb0a4888-4d06-4d75-a0c3-799f52ae4bb2",
    }
    
    var inserts []string
    
    for _, mapID := range mapIDs {
        inserts = append(inserts, f.create1v1Simulations(mapID)...)
        inserts = append(inserts, f.create4PlayerSimulations(mapID)...)
        inserts = append(inserts, f.create6PlayerSimulations(mapID)...)
    }
    
    return strings.Join(inserts, "\n")
}

func (f *SimulationFactory) create1v1Simulations(mapID string) []string {
    var inserts []string
    
    matches := [][]string{
        {"1001", "1002"},
        {"1003", "1004"},
        {"1005", "1006"},
        {"1001", "1003"},
        {"1002", "1004"},
        {"1005", "1001"},
    }
    
    for i, match := range matches {
        simData := SimulationData{
            Players:   match,
            MaxTurns:  50,
            CreatedAt: "2024-01-20T10:00:00Z",
        }
        
        jsonData, err := json.Marshal(simData)
        if err != nil {
            panic(err)
        }
        
        insert := fmt.Sprintf(`
            INSERT INTO influence.simulations (id, map_id, data, state) VALUES (
                gen_random_uuid(),
                '%s',
                '%s'::jsonb,
                'Queued'
            );`, mapID, jsonData)
        
        inserts = append(inserts, insert)
        
        if i < 3 {
            usersSim := fmt.Sprintf(`
                INSERT INTO influence.users_simulations (uid, sid, sorder) 
                SELECT uid, (SELECT id FROM influence.simulations ORDER BY queued_at DESC LIMIT 1), row_number() OVER ()
                FROM (VALUES ('%s'::bigint), ('%s'::bigint)) AS t(uid);
            `, match[0], match[1])
            
            inserts = append(inserts, usersSim)
        }
    }
    
    return inserts
}

func (f *SimulationFactory) create4PlayerSimulations(mapID string) []string {
    var inserts []string
    
    groups := [][]string{
        {"1001", "1002", "1003", "1004"},
        {"1001", "1003", "1005", "1006"},
        {"1002", "1004", "1005", "1006"},
    }
    
    for i, group := range groups {
        simData := SimulationData{
            Players:   group,
            MaxTurns:  100,
            CreatedAt: "2024-01-20T10:00:00Z",
        }
        
        jsonData, err := json.Marshal(simData)
        if err != nil {
            panic(err)
        }
        
        insert := fmt.Sprintf(`
            INSERT INTO influence.simulations (id, map_id, data, state) VALUES (
                gen_random_uuid(),
                '%s',
                '%s'::jsonb,
                'Queued'
            );`, mapID, jsonData)
        
        inserts = append(inserts, insert)
        
        if i == 0 {
            usersSim := fmt.Sprintf(`
                INSERT INTO influence.users_simulations (uid, sid, sorder) 
                SELECT uid, (SELECT id FROM influence.simulations ORDER BY queued_at DESC LIMIT 1), row_number() OVER ()
                FROM (VALUES ('%s'::bigint), ('%s'::bigint), ('%s'::bigint), ('%s'::bigint)) AS t(uid);
            `, group[0], group[1], group[2], group[3])
            
            inserts = append(inserts, usersSim)
        }
    }
    
    return inserts
}

func (f *SimulationFactory) create6PlayerSimulations(mapID string) []string {
    var inserts []string
    
    players := []string{"1001", "1002", "1003", "1004", "1005", "1006"}
    
    simData := SimulationData{
        Players:   players,
        MaxTurns:  150,
        CreatedAt: "2024-01-20T10:00:00Z",
    }
    
    jsonData, err := json.Marshal(simData)
    if err != nil {
        panic(err)
    }
    
    insert := fmt.Sprintf(`
        INSERT INTO influence.simulations (id, map_id, data, state) VALUES (
            gen_random_uuid(),
            '%s',
            '%s'::jsonb,
            'Queued'
        );`, mapID, jsonData)
    
    inserts = append(inserts, insert)
    
    usersSim := `
        INSERT INTO influence.users_simulations (uid, sid, sorder) 
        SELECT uid, (SELECT id FROM influence.simulations ORDER BY queued_at DESC LIMIT 1), row_number() OVER ()
        FROM (VALUES 
            ('1001'::bigint), 
            ('1002'::bigint), 
            ('1003'::bigint), 
            ('1004'::bigint), 
            ('1005'::bigint), 
            ('1006'::bigint)
        ) AS t(uid);`
    
    inserts = append(inserts, usersSim)
    
    return inserts
}
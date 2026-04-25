#define HAHS_SIZE 1007

typedef struct Node {
    int key;
    int val;
    struct Node* next;
} Node;

Node* table[HAHS_SIZE];

int hash(int key) {
    if(key < 0) key = -key;
    return key % HAHS_SIZE;
}

void insert(int key, int val){
    int h = hash(key);
    Node *new = malloc(sizeof(Node));

    new->key = key;
    new->val = val;
    new->next = table[h];
    table[h] = new;
}

Node* find(int key){
    int h = hash(key);
    Node* cur = table[h];

    while(cur){
        if(cur->key == key)
            return cur;

        cur = cur->next;
    }

    return NULL;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    memset(table, 0, sizeof(table));


    for(int i = 0; i < numsSize; i++){
        int need = target - nums[i];

        Node* n = find(need);

        if(n){
            int* res = malloc(2 * sizeof(int));
            res[0] = n->val + 1;
            res[1] = i + 1;
            *returnSize = 2;
            return res;
        }

        insert(nums[i], i);
    }
    return NULL;
}
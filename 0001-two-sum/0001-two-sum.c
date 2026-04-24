/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
#define SIZE 1007

typedef struct Node {
    int val;
    int key;
    struct Node* next;
} Node;

Node* table[SIZE] = {0};

int hash(int key) {
    if(key < 0) key = -key;
    return key % SIZE;
}

void insert(int key, int val) {
    int h = hash(key);
    Node *node = malloc(sizeof(Node));

    node->key = key;
    node->val = val;
    node->next = table[h];
    table[h] = node;
}

Node* find(int key){
    int h = hash(key);
    Node *cur = table[h];

    while(cur){
        if(cur->key == key) return cur;
        cur = cur->next;
    }
    return NULL;
}

int* twoSum(int* nums, int numsSize, int target, int* returnSize) {
    memset(table, 0, sizeof(table));

    for(int i = 0; i < numsSize; i++){
        int need = target -  nums[i];
        Node* n = find(need);

        if(n){
            int* res = malloc(2 * sizeof(int));
            res[0] = n->val;
            res[1] = i;
            *returnSize = 2;
            return res;
        }

        insert(nums[i], i);
    }
    *returnSize = 0;
    return NULL;

}

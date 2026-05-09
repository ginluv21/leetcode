#define HS 100100
#define HT_FOREACH(it)                          \
    for (int _i = 0; _i < HS; ++_i)             \
        for (node* it = h[_i]; it; it = it->next)


typedef struct node {
    int key;
    bool val;
    struct node* next;
} node;

node* h[HS];

int hash(int key){
    if(key < 0) key = -key;
    return key % HS;
}

bool lookup(int key){
    int a = hash(key);

    node* cur = h[a];
    while(cur){
        if(cur->key == key)
            return cur->val;

        cur = cur->next;
    }

    return false;
}

void hadd(int key, bool val){
    if(lookup(key)) return;
    int a = hash(key);
    node * new = malloc(sizeof(node));
    
    new->key = key;
    new->val = val;
    new->next = h[a];
    h[a] = new;
}

int longestConsecutive(int* nums, int numsSize) {
    memset(h, 0, sizeof(h));

    int max = 0;

    for(int i = 0; i < numsSize; i++){
        hadd(nums[i], true);
    }

    HT_FOREACH(p){

        int n = p->key;

        if(lookup(n -1))
            continue;

        int l = 1;
        int c = n;

        while(lookup(c + 1)){
            l++;
            c++;
        }

        if(max < l)
            max = l;
    }

    return max;

}
int search(int* nums, int numsSize, int target) {
    int l = 0; 
    int r = numsSize - 1;

    while(l <= r) {
        int mid = l + (r - l)/2;
        
        if(*(nums + mid) == target) 
            return mid;
        
        if (*(nums + l) <= *(nums + mid)){
            if(*(l + nums) <= target && target <= *(mid +nums))
                r = mid -1;
            else
                l = mid + 1;
        }else {
            if(*(mid + nums) <= target && target <= *(r +nums))
                l = mid + 1;
            else 
                r = mid - 1;
        }
    }

    return -1;
}
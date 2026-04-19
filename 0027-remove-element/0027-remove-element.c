int removeElement(int* nums, int numsSize, int val) {
    int s = 0;

    for(int f = 0; f < numsSize; f++){
        if(nums[f] != val){
            nums[s++] = nums[f];
        }
    }

    return s;
}
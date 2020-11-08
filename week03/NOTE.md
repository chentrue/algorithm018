分治： 
也是递归，将复杂的问题拆解成若干个简单的问题 
代码模版：
def divide_conquer(problem, param1,param2,……)
    # recursion terminator
    if problem is None:
        print_result
        return
    # prepare data
    data = prepare_data(problem)
    subproblems = split_problem(problem, data)
    # conquer subproblems
    subresult1 = self.divide_conquer(subproblems[0], p1, ……)
    subresult2 = self.divide_conquer(subproblems[1], p1, ……)
    subresult2 = self.divide_conquer(subproblems[2], p1, ……)
    ……
    #process and generate the final result   组装子结果
    result = process_result（subresult1, subresult2,subresult3……）
回溯 
思想： 
回溯法通常用最简单的递归方式实现，在反复重复上述的步骤后可能出现两种情况：
找到一个可能存在的正确答案
在尝试了所有可能的分布方法后宣告该问题没有答案



树的面试题解法一般都是递归
本质：
递归-循环
通过函数体本身进行循环
递归代码模版
def recursion(level, param1, param2)
    # recursion terminator  //递归终结条件
    if level > MAX_LEVEL:  
        process_result
        return 
    #process logic in current level   //这一层要进行的逻辑
    process(level, data……)
    #drill down
    self.recursion(level+1, p1,……) #进入下一层
    #reverse the current level status if needed 清理当前层状态
思维要点：

不要进行人肉递归(最大误区)
找到最近最简的方法，将其拆解成可重睹解决的问题（ch重复子问题）
数学归纳法思维

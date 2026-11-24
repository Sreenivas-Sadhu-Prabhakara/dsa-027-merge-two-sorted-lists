# Merge Two Sorted Linked Lists

> Striver SDE Sheet 027 · Day 5 Linked List — splice existing nodes behind a dummy head; no new nodes, no value copying. Go, stdlib only, tested.

## Problem

Merge two individually sorted linked lists into one sorted list by rearranging the existing nodes — the list-flavored half of merge sort.

## Approach

Keep a `tail` pointer to the merged list's end, starting at a local dummy node. Repeatedly splice on whichever head is smaller and advance that list. When one list runs dry, attach the other's remainder wholesale — it's already sorted and linked. The dummy head is the idiom worth stealing: it erases the "is this the first node?" special case that otherwise riddles list-building code, and `dummy.Next` is the real head at the end.

**Complexity:** O(n+m) time, O(1) extra space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).

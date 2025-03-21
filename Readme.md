# ReBAC kata

This is a kata to learn about Relation-Based Access Control (ReBAC) .
In particular, the ability to give someone permission to modify
an object temporarily.

## Subject

You are a classified ad company.
The user ask for a feature allowing to give
someone a temporary permission to edit one of their
own ads.

## Follow-up

Bonus requirements:

* Add a notion of an organization. Only the members of the organization are allowed to delegate
to other members of the organization.
* Add a super admin. The super admin can create other admins. Admins can give delegations to anyone across organisations. 

## Constraint

The permission logic code should not leak into the ads code and the user management code.

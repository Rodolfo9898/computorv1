# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    config.mk                                          :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: rperez-t <rperez-t@student.42belgium.be>   +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2026/11/16 17:00:00 by rperez-t          #+#    #+#              #
#    Updated: 2026/11/16 17:00:00 by rperez-t         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME = computor_v1
BONUS_NAME = computor_v1_bonus

# Go 1.25.4 or later (go.mod): its standard library only, nothing to download.
GO = go
GO_FLAGS = -trimpath

SRC_DIR = ./sources
# The tag the files of sources/bonus need to be compiled, as -D BONUS_MODE=1 in C.
BONUS_TAG = bonus_mode
DOCKER_IMAGE = computorv1-test
